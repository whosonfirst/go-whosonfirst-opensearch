package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
)

import (
	"context"
	_ "encoding/json"
	_ "fmt"
	"net/http"
	"log"
	"crypto/tls"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-whosonfirst-opensearch/index"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func main() {

	var os_index string
	var os_user string
	var os_pswd string

	fs := flagset.NewFlagSet("opensearch")
	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&os_user, "opensearch-user", "", "...")
	fs.StringVar(&os_pswd, "opensearch-password", "", "...")

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

	if err != nil {
		log.Fatalf("Failed to create new flagset, %v", err)
	}

	opts := &index.RunBulkIndexerOptions{
		Client: client,
		Index:  os_index,
	}

	err = index.RunBulkIndexer(ctx, opts)

	if err != nil {
		log.Fatalf("Failed to run bulk tool, %v", err)
	}
}
