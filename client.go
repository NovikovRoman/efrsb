package efrsb

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ThreePositionSwitch int

const (
	PositionSwitchUnknown ThreePositionSwitch = iota
	PositionSwitchNo
	PositionSwitchYes
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type Client struct {
	auth       *Auth
	httpClient *http.Client
	proxy      *url.URL
}

type Option func(*Client)

func WithProxy(proxy *url.URL) Option {
	return func(c *Client) {
		c.proxy = proxy
	}
}

func New(auth *Auth, opts ...Option) (c *Client) {
	c = &Client{
		auth: auth,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     10 * time.Second,
	}

	if c.proxy != nil {
		transport.Proxy = http.ProxyURL(c.proxy)
	}
	c.httpClient.Transport = transport
	return
}

func (c *Client) get(ctx context.Context, path string) (b []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.auth.Host()+path, nil)
	if err != nil {
		err = fmt.Errorf("NewRequest: %w", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.auth.token.Raw)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.httpClient.Do(req); err != nil {
		err = fmt.Errorf("Client.Do: %w", err)
		return
	}
	defer resp.Body.Close()

	if b, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("Read body: %w", err)
		return
	}

	err = responseErrHandler(b, resp.StatusCode)
	return
}
