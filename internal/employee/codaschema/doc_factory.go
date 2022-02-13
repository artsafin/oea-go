package codaschema

import (
	"github.com/artsafin/coda-go-client/codaapi"
	"sync"
)

type docFactory struct {
	docID  string
	client codaapi.ClientWithResponsesInterface
}

func NewCodaFactoryWithSharedClient(server, token, docID string, clientOpts ...codaapi.ClientOption) (*docFactory, error) {
	client, err := NewDefaultClient(server, token, clientOpts...)
	if err != nil {
		return nil, err
	}

	return &docFactory{
		docID:  docID,
		client: client,
	}, nil
}

func (f *docFactory) New() *CodaDocument {
	return &CodaDocument{
		docID:          f.docID,
		client:         f.client,
		relationsCache: &sync.Map{},
	}
}
