package main

import (
	"context"
	"log"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	wof_opensearch "github.com/whosonfirst/go-whosonfirst-opensearch"
)

func main() {

	var os_index string
	var os_username string
	var os_password string
	var os_aws_uri string
	var os_endpoint string
	var os_insecure bool

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&os_username, "opensearch-username", "", "...")
	fs.StringVar(&os_aws_uri, "opensearch-aws-uri", "", "...")
	fs.StringVar(&os_password, "opensearch-password", "", "...")
	fs.BoolVar(&os_insecure, "opensearch-insecure", false, "...")
	fs.StringVar(&os_endpoint, "opensearch-endpoint", "https://localhost:9200", "...")

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

	// START OF put me in a function in index.go

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

	// START OF put me in a function in index.go
}
