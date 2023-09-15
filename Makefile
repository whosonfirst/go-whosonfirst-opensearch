GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
