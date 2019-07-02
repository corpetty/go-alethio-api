package alethio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

// Account is the return Account structure
type Account struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Address string `json:"address"`
			Nonce   int    `json:"nonce"`
			Balance string `json:"balance"`
		} `json:"attributes"`
		Relationships struct {
			Contract struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contract"`
			Transactions struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"transactions"`
			ContractMessages struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessages"`
			EtherTransfers struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"etherTransfers"`
			TokenTransfers struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"tokenTransfers"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
	} `json:"meta"`
}

// GetAccount will return an ethereum account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}/get
func (c *Client) GetAccount(address string) (Account, error) {
	req, err := c.newRequest("GET", "/accounts/"+address, nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = c.do(req, &account)
	return account, err
}
