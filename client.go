package ecologi

import (
	"errors"
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
	}

	if host != nil {
		if *host != "" {
			c.HostURL = *host
		}
	}

	if token == nil || *token == "" {
		return nil, errors.New("no Ecologi API token provided")
	}

	c.Token = *token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) (io.ReadCloser, error) {

	var bearer string = "Bearer " + c.Token

	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	// TODO: Support idempotency keys
	// req.Header.Set("Idempotency-Key", fmt.Sprint(idempKey))

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create resource with HTTP response code: %d", res.StatusCode)
	}

	return res.Body, nil
}
