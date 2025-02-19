package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v4/opensearchutil"
)

type Indexer struct {
	client       *opensearchapi.Client
	index        string
	bulk_indexer opensearchutil.BulkIndexer
}

func NewIndexer(ctx context.Context, opensearch_client *opensearchapi.Client, opensearch_index string) (*Indexer, error) {

	idx := &Indexer{
		client: opensearch_client,
		index:  opensearch_index,
	}

	return idx, nil
}

func NewBulkIndexer(ctx context.Context, opensearch_client *opensearchapi.Client, opensearch_index string, workers int) (*Indexer, error) {

	idx, err := NewIndexer(ctx, opensearch_client, opensearch_index)

	if err != nil {
		return nil, err
	}

	bi_cfg := opensearchutil.BulkIndexerConfig{
		Index:         opensearch_index,
		Client:        opensearch_client,
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

	bulk_idx, err := opensearchutil.NewBulkIndexer(bi_cfg)

	if err != nil {
		return nil, err
	}

	idx.bulk_indexer = bulk_idx
	return idx, nil
}

func (idx *Indexer) IndexDocument(ctx context.Context, doc_id string, doc interface{}) error {

	b, err := json.Marshal(doc)

	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	return idx.IndexDocumentWithReader(ctx, doc_id, r)
}

func (idx *Indexer) IndexDocumentWithReader(ctx context.Context, doc_id string, r io.ReadSeeker) error {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	if idx.bulk_indexer == nil {

		req := opensearchapi.IndexReq{
			Index:      idx.index,
			DocumentID: doc_id,
			Body:       r,
			Params: opensearchapi.IndexParams{
				Refresh: "true",
			},
		}

		_, err := idx.client.Index(ctx, req)

		if err != nil {
			return err
		}

		return nil
	}

	bulk_item := opensearchutil.BulkIndexerItem{
		Action:     "index",
		DocumentID: doc_id,
		Body:       r,

		OnSuccess: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchapi.BulkRespItem) {
			slog.Debug("Index complete", "doc_id", doc_id)
		},

		OnFailure: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchapi.BulkRespItem, err error) {
			if err != nil {
				slog.Error("Failed to index record", "doc_id", doc_id, "error", err)
			} else {
				slog.Error("Failed to index record", "doc_id", doc_id, "type", res.Error.Type, "reason", res.Error.Reason)
			}
		},
	}

	return idx.bulk_indexer.Add(ctx, bulk_item)
}

func (idx *Indexer) Close(ctx context.Context) error {

	if idx.bulk_indexer == nil {
		return nil
	}

	return idx.bulk_indexer.Close(ctx)
}
