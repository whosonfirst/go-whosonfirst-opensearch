package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-opensearch/client"
)

func main() {

	var os_index string
	var mapping string

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&mapping, "mapping", "", "...")

	client.AppendClientFlags(fs)
	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	r, err := os.Open(mapping)

	if err != nil {
		log.Fatalf("Failed to open mappings, %v", err)
	}

	defer r.Close()

	req := opensearchapi.MappingPutReq{
		Indices: []string{
			os_index,
		},
		Body: r,
		Params: opensearchapi.MappingPutParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Indices.Mapping.Put(ctx, req)

	if err != nil {
		log.Fatalf("Failed to put mapping for '%s', %v", os_index, err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)

}
