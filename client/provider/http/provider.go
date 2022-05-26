package http

import (
	"advdiploma/client/provider"
)

var _ provider.SecretProvider = (*HTTPProvider)(nil)

type HTTPProvider struct {
	client *TokenClient
	cfg    HTTPConfig
}

func NewHTTPProvider(cfg HTTPConfig) *HTTPProvider {
	return &HTTPProvider{
		client: NewTokenClient(cfg.Timeout),
		cfg:    cfg,
	}
}
