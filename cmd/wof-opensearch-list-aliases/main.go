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

	fs := flagset.NewFlagSet("opensearch")

	client.AppendClientFlags(fs)
	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	req := opensearchapi.AliasGetReq{
		Params: opensearchapi.AliasGetParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Indices.Alias.Get(ctx, req)

	if err != nil {
		log.Fatalf("Failed to execute request, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)
}
