GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-get-mapping cmd/wof-opensearch-get-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-list-aliases cmd/wof-opensearch-list-aliases/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-indices-stats cmd/wof-opensearch-indices-stats/main.go

doc:
	{"query": { "ids": { "values": [ 1880245177 ] } } }
