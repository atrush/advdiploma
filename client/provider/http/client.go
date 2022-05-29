package http

import (
	"errors"
	"net/http"
	"time"
)

type TokenClient struct {
	http.Client
	apiToken     *string
	isAuthorised bool
}

type transport struct {
	RoundTripper http.RoundTripper
	apiToken     *string
}

func NewTokenClient(timeout time.Duration) *TokenClient {
	token := ""
	return &TokenClient{
		Client: http.Client{
			Timeout: timeout,
			Transport: &transport{
				RoundTripper: http.DefaultTransport,
				apiToken:     &token,
			},
		},
		apiToken:     &token,
		isAuthorised: false,
	}
}

func (c *TokenClient) SetToken(token string) {
	*c.apiToken = token
	c.isAuthorised = true
}

func (c *TokenClient) DropAuth() {
	*c.apiToken = ""
	c.isAuthorised = false
}

func (c *TokenClient) DoWithAuth(req *http.Request) (*http.Response, error) {
	if len(*c.apiToken) == 0 {
		return nil, errors.New("client not authorized")
	}

	req.Header.Add("Authorization", *c.apiToken)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		c.DropAuth()
	}

	return resp, err
}

// RoundTrip Implements RoundTripper interface
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {

	return t.RoundTripper.RoundTrip(req)
}
