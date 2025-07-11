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
	err := iterwriter.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to iterate, %v", err)
	}

}
