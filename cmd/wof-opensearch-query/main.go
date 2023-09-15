package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

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
		log.Fatalf("Failed to create opensearch client, %v", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{
			os_index,
		},
		Body: os.Stdin,
	}

	rsp, err := req.Do(ctx, os_client)

	if err != nil {
		log.Fatalf("Failed to perform query, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to encode response, %v", err)
	}
}
