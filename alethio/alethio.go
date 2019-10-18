package alethio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	defaultBaseURL   = "https://api.aleth.io/v1"
	defaultUserAgent = "go-alethio"

	mediaTypeV1 = "application/vnd.api+json"
)

var Opts = opts{}

// Client is the main httpClient to be built on for API calls
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	apiKey string

	client *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Alethio API
	Account          *AccountService
	Blocks           *BlocksService
	Contracts        *ContractsService
	EtherTransfers   *EtherTransfersService
	Transactions     *TransactionsService
	ContractMessages *ContractMessagesService
	LogEntries       *LogEntriesService
}

type service struct {
	client *Client
}

// NewClient creates a new Alethio API http client
func NewClient(options ...func(*Client) error) (*Client, error) {
	// defaults
	u, _ := url.Parse(defaultBaseURL)
	c := Client{
		client:    http.DefaultClient,
		BaseURL:   u,
		UserAgent: defaultUserAgent,
	}

	for _, option := range options {
		err := option(&c)
		if err != nil {
			return nil, fmt.Errorf("apply option: %s", err)
		}
	}

	c.common.client = &c
	c.Account = (*AccountService)(&c.common)
	c.Blocks = (*BlocksService)(&c.common)
	c.Contracts = (*ContractsService)(&c.common)
	c.EtherTransfers = (*EtherTransfersService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.ContractMessages = (*ContractMessagesService)(&c.common)
	c.LogEntries = (*LogEntriesService)(&c.common)

	return &c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, addedPath string, body interface{}) (*http.Request, error) {
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
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	return req, nil
}

// NewURLRequest creates a HTTP request to an URL. An absolute URL must pe specified.
// If specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewURLRequest(method, url string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = withContext(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// // If the error type is *url.Error, sanitize its URL before returning.
		// if e, ok := err.(*url.Error); ok {
		// 	if url, err := url.Parse(e.URL); err == nil {
		// 		e.URL = sanitizeURL(url).String()
		// 		return nil, e
		// 	}
		// }

		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
