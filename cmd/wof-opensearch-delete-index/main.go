package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-opensearch/v4/client"
)

func main() {

	var os_index string

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")

	client.AppendClientFlags(fs)
	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	req := opensearchapi.IndicesDeleteReq{
		Indices: []string{
			os_index,
		},
		Params: opensearchapi.IndicesDeleteParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Indices.Delete(ctx, req)

	if err != nil {
		log.Fatalf("Failed to delete index '%s', %v", os_index, err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)
}
