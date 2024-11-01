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

type AuthConfig struct {
	login string
	pass  string
	proxy *url.URL
	isDev bool
}

func NewAuthConfig(login, pass string) (c AuthConfig) {
	return AuthConfig{
		login: login,
		pass:  pass,
	}
}

func (c AuthConfig) Proxy(proxy *url.URL) AuthConfig {
	c.proxy = proxy
	return c
}

func (c AuthConfig) Dev() AuthConfig {
	c.isDev = true
	return c
}

func (c AuthConfig) Prod() AuthConfig {
	c.isDev = false
	return c
}

type Auth struct {
	httpClient *http.Client
	login      string
	pass       string
	token      *jwt.Token
	isDev      bool
}

func NewAuth(ctx context.Context, c AuthConfig) (auth *Auth, err error) {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
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
	httpClient.Transport = transport

	auth = &Auth{
		isDev:      c.isDev,
		httpClient: httpClient,
		login:      c.login,
		pass:       c.pass,
	}
	err = auth.RefreshToken(ctx)
	return
}

func (a *Auth) Host() string {
	if a.isDev {
		return devURL
	}
	return prodURL
}

func (a *Auth) IsActiveToken() (ok bool, err error) {
	var t time.Time
	if t, err = a.TokenExpirationTime(); err != nil {
		return
	}
	ok = t.After(time.Now())
	return
}

func (a *Auth) TokenExpirationTime() (t time.Time, err error) {
	var expTime *jwt.NumericDate
	if expTime, err = a.token.Claims.GetExpirationTime(); err != nil {
		err = fmt.Errorf("GetExpirationTime: %w", err)
		return
	}

	t = expTime.Time
	return
}

func (a *Auth) RefreshToken(ctx context.Context) (err error) {
	postData := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    a.login,
		Password: a.pass,
	}
	b, _ := json.Marshal(postData)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, a.Host()+"/v1/auth", bytes.NewBuffer(b))
	if err != nil {
		err = fmt.Errorf("NewRequestWithContext: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = a.httpClient.Do(req); err != nil {
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
	if a.token, _, err = parser.ParseUnverified(data.Jwt, claims); err != nil {
		err = fmt.Errorf("ParseUnverified: %w", err)
	}
	return
}
