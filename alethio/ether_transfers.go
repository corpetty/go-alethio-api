package alethio

import "context"

// EtherTransfers - ether transfer array structure for the /block/{blockhash}/etherTransfers API endpoint
type EtherTransfers struct {
	service
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			GlobalRank        []int  `json:"globalRank"`
			Total             string `json:"total"`
			TransferType      string `json:"transferType"`
			Value             string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			Block struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"block"`
			ContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"contractMessage"`
			FeeRecipient struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"feeRecipient"`
			From struct {
				Data interface{} `json:"data"`
			} `json:"from"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
			Transaction struct {
				Data interface{} `json:"data"`
			} `json:"transaction"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Links NavLinks `json:"links"`
	Meta  struct {
		Count       int `json:"count"`
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
		Page struct {
			HasNext bool `json:"hasNext"`
			HasPrev bool `json:"hasPrev"`
		} `json:"page"`
	} `json:"meta"`
}

func (et EtherTransfers) Next(ctx context.Context) (*EtherTransfers, error) {
	req, err := et.client.NewURLRequest("GET", et.Links.Next, nil)
	if err != nil {
		return nil, err
	}
	etherTransfers := new(EtherTransfers)
	_, err = et.client.Do(ctx, req, &etherTransfers)
	return etherTransfers, err
}

func (et EtherTransfers) Prev(ctx context.Context) (*EtherTransfers, error) {
	req, err := et.client.NewURLRequest("GET", et.Links.Prev, nil)
	if err != nil {
		return nil, err
	}
	etherTransfers := new(EtherTransfers)
	_, err = et.client.Do(ctx, req, &etherTransfers)
	return etherTransfers, err
}
