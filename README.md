# go-whosonfirst-opensearch
Go package for indexing Who's On First records in Opensearch.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-whosonfirst-opensearch.svg)](https://pkg.go.dev/github.com/sfomuseum/go-whosonfirst-opensearch)

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
```

### wof-opensearch-index

```
$> ./bin/wof-opensearch-index -h
  -append-spelunker-v1-properties
	Append and index auto-generated Whos On First Spelunker properties.
  -opensearch-endpoint string
    			  A fully-qualified Opensearch endpoint. (default "http://localhost:9200")
  -opensearch-index string
    		       A valid Opensearch index. (default "millsfield")
  -index-alt-files
	Index alternate geometries.
  -index-only-properties
	Only index GeoJSON Feature properties (not geometries).
  -index-spelunker-v1
	Index GeoJSON Feature properties inclusive of auto-generated Whos On First Spelunker properties.
  -iterator-uri string
    		A valid whosonfirst/go-whosonfirst-iterator/emitter URI. Supported emitter URI schemes are: directory://,featurecollection://,file://,filelist://,geojsonl://,git://,repo:// (default "repo://")
  -workers int
    	   The number of concurrent workers to index data using. Default is the value of runtime.NumCPU().
```	

For example:

```
$> bin/wof-opensearch-index \
	-index-spelunker-v1
	-opensearch-index whosonfirst \
	/usr/local/data/whosonfirst-data-admin-ca
```

### Known-knowns

#### index-spelunker-v1

* Support for generating `date:` properties derived from `edtf:` property values is currently not available. This is currently blocked on the lack of a Go language `Extended DateTime Format` parser.

## Opensearch

This code assumes Opensearch 2.x

## See also

* https://github.com/whosonfirst/go-whosonfirst-elasticsearch
* https://github.com/whosonfirst/go-whosonfirst-iterate
* https://github.com/whosonfirst/go-whosonfirst-iterate-git