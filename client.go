package nanocoinrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	VERSION = "dirty"
)

// Client represents a client connection to a collins server. Requests to the
// various APIs are done by calling functions on the various services.
type Client struct {
	Client  *http.Client
	BaseURL *url.URL

	Nano *NanoService
}

func NewClient(baseurl string) (*Client, error) {
	u, err := url.Parse(baseurl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client:  &http.Client{},
		BaseURL: u,
	}

	c.Nano = &NanoService{client: c}

	return c, nil
}

func (c *Client) NewRequest(data []byte) (*http.Request, error) {
	url := c.BaseURL.String()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go-nanocoinrpc-client "+VERSION)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, i interface{}) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, i)
	if err != nil {
		return nil, errors.New("Failed to unmarshal: " + err.Error())
	}

	return resp, nil
}
