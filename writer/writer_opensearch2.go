package writer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/url"
	"strconv"
	"sync"
	"time"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v2/opensearchutil"
	"github.com/whosonfirst/go-whosonfirst-elasticsearch/document"
	"github.com/whosonfirst/go-whosonfirst-feature/properties"
	wof_client "github.com/whosonfirst/go-whosonfirst-opensearch/client"
	wof_writer "github.com/whosonfirst/go-writer/v3"
)

func init() {
	ctx := context.Background()
	wof_writer.RegisterWriter(ctx, "opensearch", NewOpensearchV2Writer)
	wof_writer.RegisterWriter(ctx, "opensearch2", NewOpensearchV2Writer)
}

// OpensearchV2Writer is a struct that implements the `Writer` interface for writing documents to an Opensearch
// index using the github.com/opensearch-project/opensearch-go/v2 package.
type OpensearchV2Writer struct {
	wof_writer.Writer // whosonfirst/go-writer.Writer interface
	DocumentWriter    // whosonfirst/go-whosonfirst-opensearch/writer.DocumentWriter interface
	client            *opensearch.Client
	index             string
	indexer           opensearchutil.BulkIndexer
	index_alt_files   bool
	prepare_funcs     []document.PrepareDocumentFunc
	logger            *log.Logger
	waitGroup         *sync.WaitGroup
}

// NewOpensearchV2Writer returns a new `OpensearchV2Writer` instance for writing documents to an
// Opensearch index using the github.com/opensearch-project/opensearch-go/v2 package configured by 'uri' which
// is expected to take the form of:
//
//	opensearch://{HOST}:{PORT}/{INDEX}?{QUERY_PARAMETERS}
//	opensearch2://{HOST}:{PORT}/{INDEX}?{QUERY_PARAMETERS}
//
// Where {QUERY_PARAMETERS} may be one or more of the following:
// * ?bulk-index={BOOLEAN}. If true then writes will be performed using a "bulk indexer". Default is true.
// * ?workers={INT}. The number of users to enable for bulk indexing. Default is 10.
func NewOpensearchV2Writer(ctx context.Context, uri string) (wof_writer.Writer, error) {

	os_client_opts, err := wof_client.ClientOptionsFromURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create client options, %w", err)
	}

	os_client, err := wof_client.NewClientFromOptions(ctx, os_client_opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to create ES client, %w", err)
	}

	opensearch_index := os_client_opts.Index

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	logger := log.New(io.Discard, "", 0)

	wg := new(sync.WaitGroup)

	wr := &OpensearchV2Writer{
		client:    os_client,
		index:     opensearch_index,
		logger:    logger,
		waitGroup: wg,
	}

	bulk_index := true

	q_bulk_index := q.Get("bulk-index")

	if q_bulk_index != "" {

		v, err := strconv.ParseBool(q_bulk_index)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?bulk-index= parameter, %w", err)
		}

		bulk_index = v
	}

	if bulk_index {

		workers := 10

		q_workers := q.Get("workers")

		if q_workers != "" {

			w, err := strconv.Atoi(q_workers)

			if err != nil {
				return nil, fmt.Errorf("Failed to parse ?workers= parameter, %w", err)
			}

			workers = w
		}

		bi_cfg := opensearchutil.BulkIndexerConfig{
			Index:         opensearch_index,
			Client:        os_client,
			NumWorkers:    workers,
			FlushInterval: 30 * time.Second,
			OnError: func(context.Context, error) {
				if err != nil {
					slog.Error("Bulk indexer reported an error", "error", err)
				}
			},
			// OnFlushStart func(context.Context) context.Context // Called when the flush starts.
			OnFlushEnd: func(context.Context) {
				slog.Debug("Bulk indexer flush end")
			},
		}

		bi, err := opensearchutil.NewBulkIndexer(bi_cfg)

		if err != nil {
			return nil, fmt.Errorf("Failed to create bulk indexer, %w", err)
		}

		wr.indexer = bi
	}

	str_index_alt := q.Get("index-alt-files")

	if str_index_alt != "" {
		index_alt_files, err := strconv.ParseBool(str_index_alt)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?index-alt-files parameter, %w", err)
		}

		wr.index_alt_files = index_alt_files
	}

	prepare_funcs := make([]document.PrepareDocumentFunc, 0)

	prepare_funcs = append(prepare_funcs, document.PrepareSpelunkerV1Document)

	wr.prepare_funcs = prepare_funcs

	return wr, nil
}

// Write copies the content of 'fh' to the Opensearch index defined in `NewOpensearchV2Writer`.
func (wr *OpensearchV2Writer) Write(ctx context.Context, path string, r io.ReadSeeker) (int64, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return 0, fmt.Errorf("Failed to read body for %s, %w", path, err)
	}

	id, err := properties.Id(body)

	if err != nil {
		return 0, fmt.Errorf("Failed to derive ID for %s, %w", path, err)
	}

	doc_id := strconv.FormatInt(id, 10)

	alt_label, err := properties.AltLabel(body)

	if err != nil {
		return 0, fmt.Errorf("Failed to derive alt label for %s, %w", path, err)
	}

	if alt_label != "" {

		if !wr.index_alt_files {
			return 0, nil
		}

		doc_id = fmt.Sprintf("%s-%s", doc_id, alt_label)
	}

	// START OF manipulate body here...

	for _, f := range wr.prepare_funcs {

		new_body, err := f(ctx, body)

		if err != nil {
			return 0, fmt.Errorf("Failed to execute prepare func, %w", err)
		}

		body = new_body
	}

	// END OF manipulate body here...

	// Nothing to store

	if len(body) == 0 {
		slog.Debug("Document yields an empty body after prepping, skipping", "path", path)
		return 0, nil
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal %s, %v", path, err)
	}

	enc_f, err := json.Marshal(f)

	if err != nil {
		return 0, fmt.Errorf("Failed to marshal %s, %v", path, err)
	}

	// Do NOT bulk index. For example if you are using this in concert with
	// go-writer.MultiWriter running in async mode in a Lambda function where
	// the likelihood of that code being re-used across invocations is high.
	// The problem is that the first invocation will call wr.indexer.Close()
	// but then the second invocation, using the same code, will call wr.indexer.Add()
	// which will trigger a panic because the code (in opensearchutil) will try to send
	// data on a closed channel. Computers...

	if wr.indexer == nil {

		wr.waitGroup.Add(1)
		defer wr.waitGroup.Done()

		req := opensearchapi.IndexRequest{
			Index:      wr.index,
			DocumentID: doc_id,
			Body:       bytes.NewReader(enc_f),
			Refresh:    "true",
		}

		rsp, err := req.Do(ctx, wr.client)

		if err != nil {
			return 0, fmt.Errorf("Error getting response: %w", err)
		}

		defer rsp.Body.Close()

		if rsp.IsError() {
			return 0, fmt.Errorf("Failed to index document, %v", rsp.Status())
		}

		return 0, nil
	}

	// Do bulk index

	wr.waitGroup.Add(1)

	bulk_item := opensearchutil.BulkIndexerItem{
		Action:     "index",
		DocumentID: doc_id,
		Body:       bytes.NewReader(enc_f),

		OnSuccess: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchutil.BulkIndexerResponseItem) {
			slog.Debug("Index complete", "path", path, "doc_id", doc_id)
			wr.waitGroup.Done()
		},

		OnFailure: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchutil.BulkIndexerResponseItem, err error) {
			if err != nil {
				slog.Error("Failed to index record", "path", path, "error", err)
			} else {
				slog.Error("Failed to index record", "path", path, "type", res.Error.Type, "reason", res.Error.Reason)
			}

			wr.waitGroup.Done()
		},
	}

	err = wr.indexer.Add(ctx, bulk_item)

	if err != nil {
		return 0, fmt.Errorf("Failed to add bulk item for %s, %w", path, err)
	}

	return 0, nil
}

// WriterURI returns 'uri' unchanged
func (wr *OpensearchV2Writer) WriterURI(ctx context.Context, uri string) string {
	return uri
}

// Close waits for all pending writes to complete and closes the underlying writer mechanism.
func (wr *OpensearchV2Writer) Close(ctx context.Context) error {

	// Do NOT bulk index

	if wr.indexer == nil {
		wr.waitGroup.Wait()
		return nil
	}

	// Do bulk index

	err := wr.indexer.Close(ctx)

	if err != nil {
		return fmt.Errorf("Failed to close indexer, %w", err)
	}

	wr.waitGroup.Wait()

	stats := wr.indexer.Stats()

	if stats.NumFailed > 0 {
		return fmt.Errorf("Indexed (%d) documents with (%d) errors", stats.NumFlushed, stats.NumFailed)
	}

	slog.Info("Index complete", "indexed", stats.NumFlushed)
	return nil
}

// Flush() does nothing in a `OpensearchV2Writer` context.
func (wr *OpensearchV2Writer) Flush(ctx context.Context) error {
	return nil
}

// SetLogger assigns 'logger' to 'wr'.
func (wr *OpensearchV2Writer) SetLogger(ctx context.Context, logger *log.Logger) error {
	wr.logger = logger
	return nil
}

// AppendPrepareFunc will append 'fn' to the list of `go-whosonfirst-elasticsearch/document.PrepareDocumentFunc` functions
// to be applied to each document written.
func (wr *OpensearchV2Writer) AppendPrepareFunc(ctx context.Context, fn document.PrepareDocumentFunc) error {
	wr.prepare_funcs = append(wr.prepare_funcs, fn)
	return nil
}
