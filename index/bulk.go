package index

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	_ "log"
	"strconv"

	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	opensearchutil "github.com/opensearch-project/opensearch-go/v2/opensearchutil"
	"github.com/sfomuseum/go-flags/lookup"
	"github.com/sfomuseum/go-whosonfirst-opensearch/document"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
)

const FLAG_OS_ENDPOINT string = "opensearch-endpoint"
const FLAG_OS_INDEX string = "opensearch-index"
const FLAG_ITERATOR_URI string = "iterator-uri"
const FLAG_INDEX_ALT string = "index-alt-files"
const FLAG_INDEX_PROPS string = "index-only-properties"
const FLAG_INDEX_SPELUNKER_V1 string = "index-spelunker-v1"
const FLAG_APPEND_SPELUNKER_V1 string = "append-spelunker-v1-properties"
const FLAG_WORKERS string = "workers"

// type RunBulkIndexerOptions contains runtime configurations for bulk indexing
type RunBulkIndexerOptions struct {
	// BulkIndexer is a `opensearch.Transport` instance
	Client opensearchapi.Transport
	Index  string
	// PrepareFuncs are one or more `document.PrepareDocumentFunc` used to transform a document before indexing
	PrepareFuncs []document.PrepareDocumentFunc
	// IteratorURI is a valid `whosonfirst/go-whosonfirst-iterate/v2` URI string.
	IteratorURI string
	// IteratorPaths are one or more valid `whosonfirst/go-whosonfirst-iterate/v2` paths to iterate over
	IteratorPaths []string
	// IndexAltFiles is a boolean value indicating whether or not to index "alternate geometry" files
	IndexAltFiles bool
}

// PrepareFuncsFromFlagSet returns a list of zero or more known `document.PrepareDocumentFunc` functions
// based on the values in 'fs'.
func PrepareFuncsFromFlagSet(ctx context.Context, fs *flag.FlagSet) ([]document.PrepareDocumentFunc, error) {

	index_spelunker_v1, err := lookup.BoolVar(fs, FLAG_INDEX_SPELUNKER_V1)

	if err != nil {
		return nil, err
	}

	append_spelunker_v1, err := lookup.BoolVar(fs, FLAG_APPEND_SPELUNKER_V1)

	if err != nil {
		return nil, err
	}

	index_only_props, err := lookup.BoolVar(fs, FLAG_INDEX_PROPS)

	if err != nil {
		return nil, err
	}

	if index_spelunker_v1 {

		if index_only_props {
			msg := fmt.Sprintf("-%s can not be used when -%s is enabled", FLAG_INDEX_PROPS, FLAG_INDEX_SPELUNKER_V1)
			return nil, errors.New(msg)
		}

		if append_spelunker_v1 {
			msg := fmt.Sprintf("-%s can not be used when -%s is enabled", FLAG_APPEND_SPELUNKER_V1, FLAG_INDEX_SPELUNKER_V1)
			return nil, errors.New(msg)
		}
	}

	prepare_funcs := make([]document.PrepareDocumentFunc, 0)

	if index_spelunker_v1 {
		prepare_funcs = append(prepare_funcs, document.PrepareSpelunkerV1Document)
	}

	if index_only_props {
		prepare_funcs = append(prepare_funcs, document.ExtractProperties)
	}

	if append_spelunker_v1 {
		prepare_funcs = append(prepare_funcs, document.AppendSpelunkerV1Properties)
	}

	return prepare_funcs, nil
}

// RunBulkIndexer will "bulk" index a set of Who's On First documents with configuration details defined in 'opts'.
func RunBulkIndexer(ctx context.Context, opts *RunBulkIndexerOptions) error {

	os_client := opts.Client
	os_index := opts.Index

	prepare_funcs := opts.PrepareFuncs
	iterator_uri := opts.IteratorURI
	iterator_paths := opts.IteratorPaths
	index_alt := opts.IndexAltFiles

	mu := new(sync.RWMutex)

	docs := make([]interface{})
	max_docs := 1000

	iter_cb := func(ctx context.Context, path string, fh io.ReadSeeker, args ...interface{}) error {

		body, err := io.ReadAll(fh)

		if err != nil {
			return err
		}

		id_rsp := gjson.GetBytes(body, "properties.wof:id")

		if !id_rsp.Exists() {
			msg := fmt.Sprintf("%s is missing properties.wof:id", path)
			return errors.New(msg)
		}

		wof_id := id_rsp.Int()
		doc_id := strconv.FormatInt(wof_id, 10)

		alt_rsp := gjson.GetBytes(body, "properties.src:alt_label")

		if alt_rsp.Exists() {

			if !index_alt {
				return nil
			}

			doc_id = fmt.Sprintf("%s-%s", doc_id, alt_rsp.String())
		}

		// START OF manipulate body here...

		for _, f := range prepare_funcs {

			new_body, err := f(ctx, body)

			if err != nil {
				return err
			}

			body = new_body
		}

		// END OF manipulate body here...

		var f interface{}
		err = json.Unmarshal(body, &f)

		if err != nil {
			msg := fmt.Sprintf("Failed to unmarshal %s, %v", path, err)
			return errors.New(msg)
		}

		mu.Lock()
		defer mu.Unlock()

		docs = append(docs, f)

		if len(docs) < max_docs {
			return nil
		}

		// https://forum.opensearch.org/t/opensearch-go-bulk-request/9190
		// https://github.com/opensearch-project/opensearch-go/blob/1.1/opensearchutil/bulk_indexer.go

		/*
		req := opensearchapi.IndexRequest{
			Index:      os_index,
			DocumentID: doc_id,
			Body:       opensearchutil.NewJSONReader(&f),
		}

		_, err = req.Do(ctx, os_client)

		if err != nil {
			return fmt.Errorf("Failed to index %d, %w", doc_id, err)
		}
		*/

		docs = make([]interface{}, 0)

		// log.Printf("Indexed %s\n", path)
		return nil
	}

	iter, err := iterator.NewIterator(ctx, iterator_uri, iter_cb)

	if err != nil {
		return err
	}

	err = iter.IterateURIs(ctx, iterator_paths...)

	if err != nil {
		return err
	}

	return nil
}
