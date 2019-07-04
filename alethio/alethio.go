package alethio

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

const (
	defaultBaseURL = "https://api.aleth.io/v1"
)

// Client is a basic Client struct
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// NewClient creates a new httpClient and sets Alethio baseUrl
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{httpClient: httpClient, BaseURL: baseURL}

	// c.ChatService = &ChatService{client: c}
	// c.ChannelService = &ChannelService{client: c}
	// c.UserService = &UserService{client: c}
	return c
}

func (c *Client) newRequest(method, addedPath string, body interface{}) (*http.Request, error) {
	// Need to account for the 'v1' in the path, so add request path appropriately.
	u := c.BaseURL
	u.Path = path.Join(u.Path, addedPath)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Add("Authorization", "Basic "+os.Getenv("ALETHIO_APIKEY"))
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
