package common

import (
	"net/http"
	"net/url"

	"github.com/phouse512/go-coda"
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
