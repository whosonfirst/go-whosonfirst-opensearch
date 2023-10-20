GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/sign-request cmd/sign-request/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
