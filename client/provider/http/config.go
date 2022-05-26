package http

import "time"

type HTTPConfig struct {
	BaseURL     string
	AuthURL     string
	RegisterURL string

	SecretURL string

	Timeout time.Duration
}
