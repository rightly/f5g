package internal

import (
"bytes"
	"crypto/tls"
	"encoding/json"
"encoding/xml"
"fmt"
"io/ioutil"
"net/http"
"time"
)

type HttpClient struct {
	client   *http.Client
	request  *http.Request
	response *http.Response
	Body     []byte
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 1 * time.Hour,
		},
	}
}

func (c *HttpClient) SetTimeout(duration time.Duration) *HttpClient {
	c.client.Timeout = duration
	return c
}

func (c *HttpClient) DisableTLSVerify() *HttpClient {
	c.client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return c
}

func (c *HttpClient) SetRequest(method, uri string, body []byte) error {
	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return err
	}

	c.Body = []byte{}
	c.request = req

	return err
}

func (c *HttpClient) SetHeader(key, value string) *HttpClient {
	c.request.Header.Set(key, value)
	return c
}

func (c *HttpClient) SetBasicAuth(user, passwd string) *HttpClient {
	c.request.SetBasicAuth(user, passwd)
	return c
}

func (c *HttpClient) Do() error {
	resp, err := c.client.Do(c.request)
	if err != nil {
		return err
	}
	c.response = resp
	c.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if c.StatusCode()/100 == 4 || c.StatusCode()/100 == 5 {
		return fmt.Errorf("mscHttp status is %s", c.response.Status)
	}
	defer resp.Body.Close()

	return nil
}

func (c *HttpClient) StatusCode() int {
	return c.response.StatusCode
}

func (c *HttpClient) Response() *http.Response {
	return c.response
}

func (c *HttpClient) Unmarshal(v interface{}) error {
	// XML
	if c.Body[0] == '<' {
		return c.UnmarshalXML(v)
	}
	return c.UnmarshalJSON(v)
}

func (c *HttpClient) UnmarshalJSON(v interface{}) error {
	err := json.Unmarshal(c.Body, v)
	if err != nil {
		return err
	}
	return err
}

func (c *HttpClient) UnmarshalXML(v interface{}) error {
	err := xml.Unmarshal(c.Body, v)
	if err != nil {
		return err
	}

	return err
}
