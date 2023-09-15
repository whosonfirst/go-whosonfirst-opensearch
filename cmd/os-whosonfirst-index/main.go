package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
)

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	opensearchutil "github.com/opensearch-project/opensearch-go/v2/opensearchutil"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	wof_opensearch "github.com/sfomuseum/go-whosonfirst-opensearch"
	"github.com/sfomuseum/go-whosonfirst-opensearch/document"
	"github.com/sfomuseum/go-whosonfirst-opensearch/index"
)

func main() {

	var os_index string
	var os_username string
	var os_password string
	var os_aws_uri string
	var os_endpoint string
	var os_insecure bool

	var iterator_uri string
	var iterator_paths multi.MultiString

	var index_alt_files bool
	var workers int

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&os_username, "opensearch-username", "", "...")
	fs.StringVar(&os_aws_uri, "opensearch-aws-uri", "", "...")
	fs.StringVar(&os_password, "opensearch-password", "", "...")
	fs.BoolVar(&os_insecure, "opensearch-insecure", false, "...")
	fs.StringVar(&os_endpoint, "opensearch-endpoint", "https://localhost:9200", "...")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "...")
	fs.Var(&iterator_paths, "iterator-path", "...")

	fs.IntVar(&workers, "workers", 10, "...")
	fs.BoolVar(&index_alt_files, "index-alt-files", false, "...")

	flagset.Parse(fs)

	ctx := context.Background()

	os_client_opts := &wof_opensearch.ClientOptions{
		Addresses: []string{
			os_endpoint,
		},
		Insecure:          os_insecure,
		Username:          os_username,
		Password:          os_password,
		AWSCredentialsURI: os_aws_uri,
	}

	os_client, err := wof_opensearch.NewClient(ctx, os_client_opts)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	// create index here...

	settings := strings.NewReader(`{
    'settings': {
        'index': {
            'number_of_shards': 1,
            'number_of_replicas': 0
            }
        }
    }`)

	req := opensearchapi.IndicesCreateRequest{
		Index: os_index,
		Body:  settings,
	}

	_, err = req.Do(context.Background(), os_client)

	if err != nil {
		log.Fatalf("Failed to create index '%s' w/ %s: %v", os_index, os_endpoint, err)
	}

	prepare_funcs := make([]document.PrepareDocumentFunc, 0)
	prepare_funcs = append(prepare_funcs, document.PrepareSpelunkerV1Document)

	bi_cfg := opensearchutil.BulkIndexerConfig{
		Index:         os_index,
		Client:        os_client,
		NumWorkers:    workers,
		FlushInterval: 30 * time.Second,
	}

	bi, err := opensearchutil.NewBulkIndexer(bi_cfg)

	if err != nil {
		log.Fatalf("Failed to create bulk indexer, %v", err)
	}

	opts := &index.RunBulkIndexerOptions{
		BulkIndexer:   bi,
		Index:         os_index,
		IteratorURI:   iterator_uri,
		IteratorPaths: iterator_paths,
		IndexAltFiles: index_alt_files,
		PrepareFuncs:  prepare_funcs,
	}

	stats, err := index.RunBulkIndexer(ctx, opts)

	if err != nil {
		log.Fatalf("Failed to run bulk tool, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(stats)

	if err != nil {
		log.Fatalf("Failed to encode stats, %v", err)
	}
}
