package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-opensearch/v4/client"
)

func main() {

	var raw bool

	fs := flagset.NewFlagSet("opensearch")
	fs.BoolVar(&raw, "raw", false, "Display raw response JSON.")

	client.AppendClientFlags(fs)
	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	req := &opensearchapi.CatIndicesReq{}

	rsp, err := os_client.Cat.Indices(ctx, req)

	if err != nil {
		log.Fatalf("Failed to execute request, %v", err)
	}

	if raw {
		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(rsp)

		if err != nil {
			log.Fatalf("Failed to copy response, %v", err)
		}

		return
	}

	for _, idx := range rsp.Indices {
		fmt.Printf("%s\t%s\t%s\n", idx.Health, idx.UUID, idx.Index)
	}
}
