package common

import (
	"context"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/artsafin/goda"
)

// goda.ClientWithResponsesInterface
type CodaDocument struct {
	ctx   context.Context
	gc    goda.ClientWithResponsesInterface
	docId goda.DocId
}

func NewCodaDocument(baseUri, apiToken, docId string) (*CodaDocument, error) {
	bearerTokenProvider, _ := securityprovider.NewSecurityProviderBearerToken(apiToken)

	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   30 * time.Second,
	}

	godaClient, err := goda.NewClientWithResponses(
		baseUri,
		goda.WithHTTPClient(httpClient),
		goda.WithRequestEditorFn(bearerTokenProvider.Intercept),
	)

	if err != nil {
		return nil, err
	}

	return &CodaDocument{
		ctx:   context.Background(),
		gc:    godaClient,
		docId: goda.DocId(docId),
	}, nil
}

func safeSubstr(s string, from, to uint) string {
	slen := uint(len(s))
	if slen == 0 || from >= slen || to <= from {
		return ""
	}
	if to > slen {
		to = slen
	}

	rs := []rune(s)
	return string(rs[from:to])
}

func makeStatusErrorMsg(body []byte, status int) error {
	return errors.Errorf("unexpected status %d: %q", status, safeSubstr(string(body), 0, 20))
}

func (cd *CodaDocument) GetFormula(name string) (*goda.Formula, error) {
	res, err := cd.gc.GetFormulaWithResponse(cd.ctx, cd.docId, goda.FormulaIdOrName(name))
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, makeStatusErrorMsg(res.Body, res.StatusCode())
	}

	return res.JSON200, err
}

func (cd *CodaDocument) ListRows(table string, params *goda.ListRowsParams) (*goda.RowList, error) {
	res, err := cd.gc.ListRowsWithResponse(cd.ctx, cd.docId, goda.TableIdOrName(table), params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, makeStatusErrorMsg(res.Body, res.StatusCode())
	}

	return res.JSON200, err
}
