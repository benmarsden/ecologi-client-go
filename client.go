package ecologi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default Ecologi API URL
const HostURL string = "https://public.ecologi.com"

// Client -
type Client struct {
	HostURL    string
	Token      string
	HTTPClient *http.Client
}

// NewClient -
func NewClient(host *string, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default ecologi URL
		HostURL: HostURL,
		Token:   defaultDummyAPIToken,
	}

	if host != nil {
		if *host != "" {
			c.HostURL = *host
		}
	}
	if token != nil {
		if *token != "" {
			c.Token = *token
		}
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) (io.ReadCloser, error) {

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if req.Method == http.MethodPost {
		if res.StatusCode != http.StatusCreated {
			return nil, fmt.Errorf("failed to create resource with HTTP response code: %d", res.StatusCode)
		}
	}

	if req.Method == http.MethodGet {
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get resource with HTTP response code: %d", res.StatusCode)
		}
	}

	return res.Body, nil
}
