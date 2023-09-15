package opensearch

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/aaronland/go-aws-auth"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

type ClientOptions struct {
	Addresses         []string
	Insecure          bool
	Username          string
	Password          string
	AWSCredentialsURI string
}

func NewClient(ctx context.Context, opts *ClientOptions) (*opensearch.Client, error) {

	os_cfg := opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: opts.Insecure,
			},
		},
		Addresses: opts.Addresses,
	}

	if opts.AWSCredentialsURI != "" {

		aws_cfg, err := auth.NewConfig(ctx, opts.AWSCredentialsURI)

		if err != nil {
			return nil, fmt.Errorf("Failed to create new AWS config, %w", err)
		}

		signer, err := requestsigner.NewSignerWithService(aws_cfg, "es")

		if err != nil {
			return nil, fmt.Errorf("Failed to create request signer, %w", err)
		}

		os_cfg.Signer = signer

	} else {

		os_cfg.Username = opts.Username
		os_cfg.Password = opts.Password
	}

	client, err := opensearch.NewClient(os_cfg)

	if err != nil {
		return nil, fmt.Errorf("Failed to create client, %w", err)
	}

	return client, nil
}
