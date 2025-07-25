package main

// Everything in here needs to be refactored because the /opensearchapi/api-aliases.go code
// doesn't actually return a list of aliases...

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

	req := opensearchapi.AliasesReq{
		Params: opensearchapi.AliasesParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Aliases(ctx, req)

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
