package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
	_ "github.com/whosonfirst/go-whosonfirst-iterate-organization"
	_ "github.com/whosonfirst/go-whosonfirst-opensearch/v4/writer"
)

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-iterwriter/application/iterwriter"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := iterwriter.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to iterate, %v", err)
	}

}
