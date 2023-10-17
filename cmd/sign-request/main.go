package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aaronland/go-aws-auth"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

func main() {

	var method string
	var uri string

	var credentials_uri string

	flag.StringVar(&method, "method", "GET", "")
	flag.StringVar(&uri, "uri", "", "")
	flag.StringVar(&credentials_uri, "credentials-uri", "", "")

	flag.Parse()

	body_r := strings.NewReader(strings.Join(flag.Args(), " "))

	ctx := context.Background()

	aws_cfg, err := auth.NewConfig(ctx, credentials_uri)

	if err != nil {
		log.Fatalf("Failed to create new AWS config, %v", err)
	}

	signer, err := requestsigner.NewSignerWithService(aws_cfg, "es")

	if err != nil {
		log.Fatalf("Failed to create request signer, %v", err)
	}

	req, err := http.NewRequest(method, uri, body_r)

	if err != nil {
		log.Fatalf("Failed to create new HTTP request, %v", err)
	}

	err = signer.SignRequest(req)

	if err != nil {
		log.Fatalf("Failed to sign request, %v", err)
	}

	err = req.Write(os.Stdout)

	if err != nil {
		log.Fatalf("Failed to write request, %v", err)
	}
}
