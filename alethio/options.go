package alethio

import (
	"net/http"
	"net/url"
)

type opts struct {
}

// APIKey is used to set the API Key
func (o opts) APIKey(k string) func(*Client) error {
	return func(c *Client) error {
		c.apiKey = k
		return nil
	}
}

// URL is used to set the base url for the API
func (o opts) URL(u string) func(*Client) error {
	return func(c *Client) error {
		var err error
		c.BaseURL, err = url.Parse(u)
		return err
	}
}

// URL is used to set the base url for the API
func (o opts) UserAgent(ua string) func(*Client) error {
	return func(c *Client) error {
		c.UserAgent = ua
		return nil
	}
}

// HTTPClient is used to set a different http client than the default one
func (o opts) HTTPClient(h *http.Client) func(*Client) error {
	return func(c *Client) error {
		c.client = h
		return nil
	}
}
