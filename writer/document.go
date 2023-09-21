package writer

import (
	"context"
	
	"github.com/whosonfirst/go-whosonfirst-elasticsearch/document"	
)

type DocumentWriter interface {
	AppendPrepareFunc(context.Context, document.PrepareDocumentFunc) error
}
