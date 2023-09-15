package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
)

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aaronland/go-aws-auth"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	opensearchutil "github.com/opensearch-project/opensearch-go/v2/opensearchutil"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-whosonfirst-opensearch/document"
	"github.com/sfomuseum/go-whosonfirst-opensearch/index"
)

func main() {

	var os_index string
	var os_user string
	var os_pswd string
	var os_aws_uri string
	var os_endpoint string
	var os_insecure bool

	var iterator_uri string
	var iterator_paths multi.MultiString

	var index_alt_files bool
	var workers int

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&os_user, "opensearch-user", "", "...")
	fs.StringVar(&os_aws_uri, "opensearch-aws-uri", "", "...")
	fs.StringVar(&os_pswd, "opensearch-password", "", "...")
	fs.BoolVar(&os_insecure, "opensearch-insecure", false, "...")
	fs.StringVar(&os_endpoint, "opensearch-endpoint", "https://localhost:9200", "...")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "...")
	fs.Var(&iterator_paths, "iterator-path", "...")

	fs.IntVar(&workers, "workers", 10, "...")
	fs.BoolVar(&index_alt_files, "index-alt-files", false, "...")

	flagset.Parse(fs)

	ctx := context.Background()

	os_cfg := opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: os_insecure,
			},
		},
		Addresses: []string{
			os_endpoint,
		},
	}

	if os_aws_uri != "" {

		aws_cfg, err := auth.NewConfig(ctx, os_aws_uri)

		if err != nil {
			log.Fatalf("Failed to create new AWS config, %v", err)
		}

		signer, err := requestsigner.NewSignerWithService(aws_cfg, "es")

		if err != nil {
			log.Fatalf("Failed to create request signer, %v", err)
		}

		os_cfg.Signer = signer

	} else {

		os_cfg.Username = os_user
		os_cfg.Password = os_pswd
	}

	client, err := opensearch.NewClient(os_cfg)

	if err != nil {
		log.Fatalf("Failed to create client, %w", err)
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

	_, err = req.Do(context.Background(), client)

	if err != nil {
		log.Fatalf("Failed to create index '%s' w/ %s: %v", os_index, os_endpoint, err)
	}

	prepare_funcs := make([]document.PrepareDocumentFunc, 0)
	prepare_funcs = append(prepare_funcs, document.PrepareSpelunkerV1Document)

	bi_cfg := opensearchutil.BulkIndexerConfig{
		Index:         os_index,
		Client:        client,
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
