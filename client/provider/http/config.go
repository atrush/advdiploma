package http

import "time"

type HTTPConfig struct {
	BaseURL     string
	AuthURL     string
	RegisterURL string
	SyncListURL string

	SecretURL string

	Timeout time.Duration
}
