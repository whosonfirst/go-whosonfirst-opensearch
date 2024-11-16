package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-opensearch/client"
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
		log.Fatalf("Failed to create opensearch client, %v", err)
	}

	q := strings.Join(fs.Args(), " ")

	req := &opensearchapi.SearchReq{
		Indices: []string{
			os_index,
		},
		Body: strings.NewReader(q),
		Params: opensearchapi.SearchParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Search(ctx, req)

	if err != nil {
		log.Fatalf("Failed to perform query, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)
}
