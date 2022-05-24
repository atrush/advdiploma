package http

import "time"

type HTTPConfig struct {
	BaseURL     string
	AuthURL     string
	RegisterURL string

	timeout time.Duration
}
