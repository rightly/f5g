package bigip

import (
	"fmt"
	"oss.navercorp.com/seunghwan.na/f5g/pkg/internal"
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

func (c *client)buildUrl(path ...string) string {
	assembled := fmt.Sprintf("%s://%s", c.scheme, c.domain)
	i := 0
	for _, v := range path {
		assembled += v
		if i < len(path) - 1 {
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
	httpClient := internal.NewHttpClient()
	if c.disableTlsVerify {
		httpClient.DisableTLSVerify()
	}
	
	err := httpClient.SetRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("iControlRequest.SetRequest fail: %s", err)
	}

	httpClient.
		SetTimeout(c.timeout).
		SetBasicAuth(c.a.user, c.a.passwd)
	
	err = httpClient.Do()
	if err != nil {
		return fmt.Errorf("iControlRequest.Do fail: %s", err)
	}
	
	err = httpClient.Unmarshal(r)
	if err != nil {
		return fmt.Errorf("iControlRequest.Unmarshal fail: %s", err)
	}

	return nil
}