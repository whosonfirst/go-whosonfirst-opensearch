package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-opensearch/client"
)

func main() {

	var os_index string
	var settings string

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&settings, "settings", "", "...")

	client.AppendClientFlags(fs)
	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	r, err := os.Open(settings)

	if err != nil {
		log.Fatalf("Failed to open settingss, %v", err)
	}

	defer r.Close()

	req := opensearchapi.IndicesPutSettingsRequest{
		Index: []string{
			os_index,
		},
		Body:   r,
		Pretty: true,
	}

	rsp, err := req.Do(context.Background(), os_client)

	if err != nil {
		log.Fatalf("Failed to put settings for '%s', %v", os_index, err)
	}

	defer rsp.Body.Close()

	_, err = io.Copy(os.Stdout, rsp.Body)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	if rsp.IsError() {
		os.Exit(1)
	}

	os.Exit(0)

}
