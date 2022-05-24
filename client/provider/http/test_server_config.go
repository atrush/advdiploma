package http

import (
	"time"
)

type serverTestConfig struct {
	returnHeaders map[string]string
	returnStatus  int
	returnBody    string
	sleep         time.Duration

	reqHeaders map[string]string
	reqBody    string
	reqURL     string
}
type serverTestParam func(*serverTestConfig)

func (c *serverTestConfig) New(opts ...serverTestParam) serverTestConfig {
	res := *c
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(&res)
		}
	}

	return res
}

func withReturnHeaders(v map[string]string) serverTestParam {
	return func(h *serverTestConfig) {
		h.returnHeaders = v
	}
}

func withReturnStatus(v int) serverTestParam {
	return func(h *serverTestConfig) {
		h.returnStatus = v
	}
}

func withReturnBody(v string) serverTestParam {
	return func(h *serverTestConfig) {
		h.returnBody = v
	}
}

func withSleep(v time.Duration) serverTestParam {
	return func(h *serverTestConfig) {
		h.sleep = v
	}
}

func withReqHeaders(v map[string]string) serverTestParam {
	return func(h *serverTestConfig) {
		h.reqHeaders = v
	}
}

func withReqURL(url string) serverTestParam {
	return func(h *serverTestConfig) {
		h.reqURL = url
	}
}

func withReqBody(url string) serverTestParam {
	return func(h *serverTestConfig) {
		h.reqURL = url
	}
}
