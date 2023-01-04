package itunes

import (
	"net/http"
)

// ClientOptionFunc can be used to customize a new iTunes Search API client.
type ClientOptionFunc func(*Client) error

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.client = httpClient
		return nil
	}
}
