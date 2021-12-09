package main

import (
	"encoding/json"
	"fmt"
	"github.com/rightly/f5g/chttp"
	"strings"
	"time"
)

type client struct {
	a                *authentication
	timeout          time.Duration
	disableTlsVerify bool
	scheme           string
	domain           string
}

type authentication struct {
	user   string
	passwd string
}

type Error struct {
	Code       int           `json:"code"`
	Message    string        `json:"message"`
	ErrorStack []interface{} `json:"errorStack"`
	APIError   int           `json:"apiError"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func newError(code int, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		ErrorStack: nil,
		APIError:   0,
	}
}

func New(scheme, domain, user, password string) *client {
	return &client{
		a: &authentication{
			user:   user,
			passwd: password,
		},
		disableTlsVerify: false,
		timeout:          10 * time.Second,
		scheme:           scheme,
		domain:           domain,
	}
}

func CommonId(name string) string {
	return fmt.Sprintf("~Common~%s", name)
}

func (c *client) buildUrl(path ...string) string {
	assembled := fmt.Sprintf("%s://%s", c.scheme, c.domain)
	i := 0
	for _, v := range path {
		assembled += v
		if i < len(path)-1 {
			assembled += "/"
		}
		i++
	}

	return assembled
}

func (c *client) SetTimeout(duration time.Duration) *client {
	c.timeout = duration

	return c
}

func (c *client) DisableTLSVerify() *client {
	c.disableTlsVerify = true
	return c
}

func (c *client) iControlRequest(method, url string, body []byte, r interface{}) error {
	httpClient := chttp.NewHttpClient()
	if c.disableTlsVerify {
		httpClient.DisableTLSVerify()
	}

	err := httpClient.SetRequest(method, url, body)
	if err != nil {
		return newError(500, fmt.Errorf("iControlRequest.SetRequest fail: %s", err).Error())
	}

	httpClient.
		SetTimeout(c.timeout).
		SetBasicAuth(c.a.user, c.a.passwd)

	err = httpClient.Do()
	if err != nil {
		errResp := newError(0, "")
		_ = httpClient.UnmarshalJSON(errResp)
		errResp.Message = "iControlRequest.Do fail: " + errResp.Message
		return errResp
	}
	if len(httpClient.Body) > 0 {
		err = httpClient.UnmarshalJSON(r)
		if err != nil {
			return newError(500, fmt.Errorf("iControlRequest.Unmarshal fail: %s", err).Error())
		}
	}

	return nil
}

func IdEncoding(id string) string {
	return strings.ReplaceAll(strings.ReplaceAll(id, blank, "%20"), slash, tilde)
}
