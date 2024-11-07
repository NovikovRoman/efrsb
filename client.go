package efrsb

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	httpClient *http.Client
	login      string
	pass       string
	proxy      *url.URL

	token *jwt.Token
	isDev bool
}

type Option func(*Client)

func Proxy(proxy *url.URL) Option {
	return func(c *Client) {
		c.proxy = proxy
	}
}

func Dev() Option {
	return func(c *Client) {
		c.isDev = true
	}
}

func Prod() Option {
	return func(c *Client) {
		c.isDev = false
	}
}

func New(login, pass string, opts ...Option) (c *Client) {
	c = &Client{
		login: login,
		pass:  pass,
		httpClient: &http.Client{
			Timeout: time.Second * 45,
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

func (c *Client) TokenExpirationTime() (t time.Time, err error) {
	if c.token == nil {
		return
	}
	var expTime *jwt.NumericDate
	if expTime, err = c.token.Claims.GetExpirationTime(); err != nil {
		err = fmt.Errorf("GetExpirationTime: %w", err)
		return
	}

	t = expTime.Time
	return
}

func (c *Client) IsActiveToken() (ok bool, err error) {
	var t time.Time
	if t, err = c.TokenExpirationTime(); err != nil {
		return
	}
	ok = t.After(time.Now())
	return
}

func (c *Client) Auth(ctx context.Context) (err error) {
	return c.RefreshToken(ctx)
}

func (c *Client) RefreshToken(ctx context.Context) (err error) {
	postData := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    c.login,
		Password: c.pass,
	}
	b, _ := json.Marshal(postData)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, c.host()+"/v1/auth", bytes.NewBuffer(b))
	if err != nil {
		err = fmt.Errorf("NewRequestWithContext: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.httpClient.Do(req); err != nil {
		err = fmt.Errorf("Client.Do: %w", err)
		return
	}
	defer resp.Body.Close()

	if b, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("ReadAll: %w", err)
	}

	if err = responseErrHandler(b, resp.StatusCode); err != nil {
		return
	}

	data := struct {
		Jwt string `json:"jwt"`
	}{}
	if err = json.Unmarshal(b, &data); err != nil {
		err = fmt.Errorf("Unmarshal token: %w Body: %s", err, b)
	}

	parser := jwt.NewParser()
	claims := make(jwt.MapClaims)
	if c.token, _, err = parser.ParseUnverified(data.Jwt, claims); err != nil {
		err = fmt.Errorf("ParseUnverified: %w", err)
	}
	return
}

func (c *Client) host() string {
	if c.isDev {
		return devURL
	}
	return prodURL
}

func (c *Client) get(ctx context.Context, path string) (b []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.host()+path, nil)
	if err != nil {
		err = fmt.Errorf("NewRequest: %w", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.token.Raw)
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
