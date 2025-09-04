GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

GOOS=darwin
GOARCH=arm64

cli:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-get-mapping cmd/wof-opensearch-get-mapping/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-settings cmd/wof-opensearch-put-settings/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-list-indices cmd/wof-opensearch-list-indices/main.go
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-indices-stats cmd/wof-opensearch-indices-stats/main.go

doc:
	{"query": { "ids": { "values": [ 1880245177 ] } } }

# Debugging targets

# This is for debugging. Do not change this at your own risk.
# (That means you should change this.)
OSPSWD=dkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj

spelunker:
	@make spelunker-mappings
	@make spelunker-settings

spelunker-mappings:
	bin/wof-opensearch-create-index \
		-opensearch-endpoint 'https://admin:$(OSPSWD)@localhost:9200' \
		-opensearch-insecure \
		-opensearch-index spelunker \
		-settings /usr/local/whosonfirst/whosonfirst-opensearch/schema/2.x/mappings.spelunker.json

spelunker-settings:
	bin/wof-opensearch-put-settings \
		-opensearch-endpoint 'https://admin:$(OSPSWD)@localhost:9200' \
		-opensearch-insecure \
		-opensearch-index spelunker \
		-settings /usr/local/whosonfirst/whosonfirst-opensearch/schema/2.x/settings.spelunker.json

index-repo:
	bin/wof-opensearch-index \
		-writer-uri 'constant://?val=opensearch2%3A%2F%2Flocalhost%3A9200%2Fspelunker%3Fusername%3Dadmin%26password%3D$(OSPSWD)%26insecure%3Dtrue%26require-tls%3Dtrue' \
		$(REPO)
