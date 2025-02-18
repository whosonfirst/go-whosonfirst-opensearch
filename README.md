# go-whosonfirst-opensearch

Go package for indexing Who's On First records in OpenSearch.

## Documentation

Documentation is incomplete at this time.

## A note about version numbers

The `github.com/whosonfirst/go-whosonfirst-opensearch` package targets the `opensearch-project/opensearch-go/v2` package.

Starting with `github.com/whosonfirst/go-whosonfirst-opensearch/v4` version numbers will track which their corresponding major version of the `opensearch-project/opensearch-go` package.

This is to account for things like a whole host of backwards incompatible changes between `opensearch-go/v2` and `opensearch-go/v3`.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
go build -mod vendor -ldflags="-s -w" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
go build -mod vendor -ldflags="-s -w" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
go build -mod vendor -ldflags="-s -w" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
go build -mod vendor -ldflags="-s -w" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
```

### wof-opensearch-index

```
$> ./bin/wof-opensearch-index \
	-writer-uri 'constant://?val=opensearch2%3A%2F%2Flocalhost%3A9200%2Fspelunker%3Frequire-tls%3Dtrue%26insecure%3Dtrue%26debug%3Dfalse%26username%3Dadmin%26password%3Ds33kret' \
	/usr/local/data/whosonfirst-data-admin-xy

2024/03/09 22:36:53 time to index paths (1) 1.530319599s
2024/03/09 22:36:57 INFO Index complete indexed=18432
```

Note the unfortunate need to URL escape the `-writer-uri=constant://?val=` parameter which unescaped is the actual `go-whosonfirst-opensearch/writer.OpensearchV2Writer` URI that takes the form of:

```
opensearch2://localhost:9200/spelunker?require-tls=true&insecure=true&debug=false&username=admin&password=s33kret
```

The `wof-opensearch-index` application however expects a [gocloud.dev/runtimevar](https://gocloud.dev/howto/runtimevar/) URI so that you don't need to deply production configuration values with sensitive values (like OpenSearch admin passwords) exposed in them. Under the hood the `wof-opensearch-index` application is using the [sfomuseum/runtimevar](https://github.com/sfomuseum/runtimevar) package to manage the details and this needs to be updated to allow plain (non-runtimevar) strings. Or maybe the `wof-opensearch-index` application needs to be updated. Either way something needs to be updated to avoid the hassle of always needing to URL-escape things.

### wof-opensearch-create-index

Create a new OpenSearch index with optional settings and mappings.

```
$> ./bin/wof-opensearch-create-index \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
	-opensearch-index collection \
	-settings /usr/local/sfomuseum/es-sfomuseum-schema/schema/7.4/mappings.collection.json

{"acknowledged":true,"shards_acknowledged":true}
```

### wof-opensearch-delete-index

Delete an OpenSearch index.

```
$> ./bin/wof-opensearch-delete-index \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
	-opensearch-index collection \
```

### wof-opensearch-index

Index one or more [whosonfirst/go-whosonfirst-iterate/v2](https://github.com/whosonfirst/go-whosonfirst-iterate) sources in an OpenSearch index.

```
$> ./bin/wof-opensearch-index \
	-writer-uri 'constant://?val=...' \
	/usr/local/data/sfomuseum-data-collection-classifications/
	
2023/09/15 23:42:19 time to index paths (1) 10.551282589s
```

#### Writer URIs

TBW. Have a look at [writer/writer_opensearch2.go](writer/writer_opensearch2.go) and [whosonfirst/go-whosonfirst-iterwriter](https://github.com/whosonfirst/go-whosonfirst-iterwriter) for the time being.

### wof-opensearch-list-aliases

List all the aliases for an OpenSearch instance.

```
$> ./bin/wof-opensearch-list-aliases \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
```

### wof-opensearch-get-mapping

Retrieve the mappings for an OpenSearch index.

```
$> ./bin/wof-opensearch-get-mapping \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
	-opensearch-index collection
```

### wof-opensearch-put-mapping

Assign and update _new_ mappings in an OpenSearch index. Remember: It is not possible to change existing mappings in an index. 

```
$> ./bin/wof-opensearch-put-mapping \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
	-opensearch-index collection \
	-mapping mapping.json
```

### wof-opensearch-query

Query an OpenSearch index.

```
$> ./bin/wof-opensearch-query \
	-opensearch-aws-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-opensearch-endpoint https://{OPENSEARCH_DOMAIN}.{REGION}.es.amazonaws.com \
	-opensearch-index collection \
	'{"query": { "ids": { "values": [ 1880245177 ] } } }'  | jq
	
{
  "took": 914,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 1,
      "relation": "eq"
    },
    "max_score": 1,
    "hits": [
      {
        "_index": "collection",
        "_id": "1880245177",
        "_score": 1,
        "_source": {
          "counts:names_languages": 0,
          "counts:names_prefered": 0,
          "counts:names_total": 0,
          "counts:names_variant": 0,
          "edtf:cessation": "",
          "edtf:inception": "",
          "geom:area": 0,
          "geom:bbox": "-122.386155,37.616357,-122.386155,37.616357",
          "geom:latitude": 37.616357,
          "geom:longitude": -122.386155,
          "mz:is_current": 1,
          "sfomuseum:category": "Scheduling / Ticketing",
          "sfomuseum:classification_id": -1,
          "sfomuseum:collection": "Aviation Museum",
          "sfomuseum:placetype": "subcategory",
          "sfomuseum:subcategory": "Luggage Tag, Crew",
          "src:geom": "sfomuseum",
          "translations": [],
          "wof:belongsto": [
            1511214277,
            102527513,
            1762679689,
            1511214203,
            102191575,
            85633793,
            102087579,
            85922583,
            85688637
          ],
          "wof:country": "US",
          "wof:created": 1693337008,
          "wof:geomhash": "30c8a918561c84bb2daa2b97fc7c5353",
          "wof:hierarchy": [
            {
              "building_id": 1511214277,
              "campus_id": 102527513,
              "category_id": 1762679689,
              "collection_id": 1511214203,
              "continent_id": 102191575,
              "country_id": 85633793,
              "county_id": 102087579,
              "locality_id": 85922583,
              "neighbourhood_id": -1,
              "region_id": 85688637,
              "subcategory_id": 1880245177,
              "wing_id": 1511214203
            }
          ],
          "wof:id": 1880245177,
          "wof:lastmodified": 1693337008,
          "wof:name": "Luggage Tag, Crew",
          "wof:parent_id": 1762679689,
          "wof:placetype": "custom",
          "wof:placetype_alt": "subcategory",
          "wof:placetype_id": 1729783759,
          "wof:placetype_names": [
            "custom",
            "subcategory"
          ],
          "wof:repo": "sfomuseum-data-collection-classfications",
          "wof:superseded_by": [],
          "wof:supersedes": []
        }
      }
    ]
  }
}
```

## See also

* https://pkg.go.dev/github.com/opensearch-project/opensearch-go
* https://github.com/whosonfirst/go-whosonfirst-elasticsearch
* https://github.com/whosonfirst/go-writer
* https://github.com/whosonfirst/go-whosonfirst-iterwriter
* https://github.com/whosonfirst/go-whosonfirst-iterate
* https://github.com/whosonfirst/go-whosonfirst-iterate-git
* https://github.com/aaronland/go-aws-auth