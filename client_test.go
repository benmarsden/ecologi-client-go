package ecologi

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	tcs := []struct {
		desc           string
		host           string
		token          string
		expectedErr    error
		expectedClient *Client
	}{
		{
			desc:        "empty string host",
			host:        "",
			token:       defaultDummyAPIToken,
			expectedErr: nil,
			expectedClient: &Client{
				HostURL:    "https://public.ecologi.com",
				Token:      defaultDummyAPIToken,
				HTTPClient: &http.Client{Timeout: 10 * time.Second},
			},
		},
		{
			desc:        "nil host",
			token:       defaultDummyAPIToken,
			expectedErr: nil,
			expectedClient: &Client{
				HostURL:    "https://public.ecologi.com",
				Token:      defaultDummyAPIToken,
				HTTPClient: &http.Client{Timeout: 10 * time.Second},
			},
		},
		{
			desc:           "empty string token",
			host:           "https://public.ecologi.com",
			token:          "",
			expectedErr:    errors.New("no Ecologi API token provided"),
			expectedClient: nil,
		},
		{
			desc:           "nil token",
			host:           "https://public.ecologi.com",
			expectedErr:    errors.New("no Ecologi API token provided"),
			expectedClient: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			c, err := NewClient(&tc.host, &tc.token)

			assert.Equal(t, tc.expectedErr, err)

			assert.Equal(t, tc.expectedClient, c)
		})
	}
}
