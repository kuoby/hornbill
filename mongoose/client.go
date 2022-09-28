package mongoose

import (
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpCli *http.Client
	baseUrl string
}

// NewClient returns a new Client.
func NewClient(baseUrl string) *Client {
	return &Client{
		httpCli: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseUrl: strings.TrimSuffix(baseUrl, "/") + "/api/hornbill",
	}
}
