package http

import (
	"net/http"
)

type TokenClient struct {
	http.Client
	apiToken *string
}

type transport struct {
	RoundTripper http.RoundTripper
	apiToken     *string
}

func NewTokenClient() *TokenClient {
	token := ""
	return &TokenClient{
		Client: http.Client{
			Transport: &transport{
				RoundTripper: http.DefaultTransport,
				apiToken:     &token,
			},
		},
		apiToken: &token,
	}
}

func (c *TokenClient) SetToken(token string) {
	*c.apiToken = token
}

// RoundTrip Implements RoundTripper interface
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(*t.apiToken) > 0 {
		req.Header.Add("Authorization", *t.apiToken)
	}

	return t.RoundTripper.RoundTrip(req)
}
