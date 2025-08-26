package functionaloption

import (
	"net/http"
	"time"
)

type HttpClient struct {
	Timeout time.Duration
}

func DefaultHttpClient() HttpClient {
	return HttpClient{
		Timeout: 10 * time.Second,
	}
}

type Option func(*HttpClient)

func WithTimeout(timeout time.Duration) Option {
	return func(hc *HttpClient) {
		hc.Timeout = timeout
	}
}

func NewHttpClient(opts ...Option) *http.Client {
	httpClient := DefaultHttpClient()
	for _, opt := range opts {
		opt(&httpClient)
	}
	return &http.Client{
		Timeout: httpClient.Timeout,
	}
}
