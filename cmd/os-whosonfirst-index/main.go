package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
)

import (
	"context"
	"crypto/tls"
	_ "encoding/json"
	_ "fmt"
	"log"
	"net/http"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-whosonfirst-opensearch/document"
	"github.com/sfomuseum/go-whosonfirst-opensearch/index"
)

func main() {

	var os_index string
	var os_user string
	var os_pswd string
	var iterator_uri string
	var iterator_paths multi.MultiString

	var index_alt_files bool

	fs := flagset.NewFlagSet("opensearch")
	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&os_user, "opensearch-user", "", "...")
	fs.StringVar(&os_pswd, "opensearch-password", "", "...")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "...")
	fs.Var(&iterator_paths, "iterator-path", "...")

	fs.BoolVar(&index_alt_files, "index-alt-files", false, "...")

	flagset.Parse(fs)

	ctx := context.Background()

	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
		},
		Addresses: []string{"https://localhost:9202"},
		Username:  os_user,
		Password:  os_pswd,
	})

	if err != nil {
		log.Fatalf("Failed to create client, %w", err)
	}

	// create index here...

	prepare_funcs := make([]document.PrepareDocumentFunc, 0)
	prepare_funcs = append(prepare_funcs, document.AppendSpelunkerV1Properties)

	opts := &index.RunBulkIndexerOptions{
		Client:        client,
		Index:         os_index,
		IteratorURI:   iterator_uri,
		IteratorPaths: iterator_paths,
		IndexAltFiles: index_alt_files,
		PrepareFuncs:  prepare_funcs,
	}

	err = index.RunBulkIndexer(ctx, opts)

	if err != nil {
		log.Fatalf("Failed to run bulk tool, %v", err)
	}
}
