module github.com/sfomuseum/go-whosonfirst-opensearch

go 1.21

// Note that elastic/go-elasticsearch/v7 v7.13.0 is the last version known to work with AWS
// Elasticsearch instances. v7.14.0 and higher will fail with this error message:
// "the client noticed that the server is not a supported distribution of Elasticsearch"
// Good times...

require (
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/opensearch-project/opensearch-go/v2 v2.3.0
	github.com/sfomuseum/go-edtf v0.3.1
	github.com/sfomuseum/go-flags v0.8.2
	github.com/tidwall/gjson v1.14.0
	github.com/tidwall/sjson v1.2.4
	github.com/whosonfirst/go-whosonfirst-iterate-git/v2 v2.1.0
	github.com/whosonfirst/go-whosonfirst-iterate/v2 v2.0.1
	github.com/whosonfirst/go-whosonfirst-placetypes v0.3.0
)

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/aaronland/go-json-query v0.1.1 // indirect
	github.com/aaronland/go-roster v0.0.2 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/whosonfirst/go-ioutil v1.0.1 // indirect
	github.com/whosonfirst/go-whosonfirst-crawl v0.2.1 // indirect
	github.com/whosonfirst/walk v0.0.1 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
