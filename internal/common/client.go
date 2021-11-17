package common

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/artsafin/go-coda"
)

type AuthTransport struct {
	Transport http.RoundTripper
	apiToken  string
}

func (adt *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+adt.apiToken)
	return adt.Transport.RoundTrip(req)
}

type CodaClient struct {
	coda.Client
}

func NewCodaClient(baseUri, apiToken string) *CodaClient {
	transport := &AuthTransport{Transport: http.DefaultTransport, apiToken: apiToken}
	http := &http.Client{
		Transport: transport,
	}

	url, err := url.Parse(baseUri)
	if err != nil {
		panic(err)
	}

	return &CodaClient{
		Client: coda.Client{
			BaseURL:    url,
			UserAgent:  "oea-go",
			HttpClient: http,
		},
	}
}

func Query(k, v string) string {
	k = strings.ReplaceAll(k, "\"", "\\\"")
	v = strings.ReplaceAll(v, "\"", "\\\"")

	return fmt.Sprintf("\"%s\":\"%s\"", k, v)
}

type QueryParam func(p *coda.ListRowsParameters)

type QueryViewParam func(p *coda.ListViewRowsParameters)
