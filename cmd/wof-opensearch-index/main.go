package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v3"
	_ "github.com/whosonfirst/go-whosonfirst-opensearch/v4/writer"

	"github.com/whosonfirst/go-whosonfirst-iterwriter/v4/app/iterwriter"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := iterwriter.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to iterate, %v", err)
	}

}
