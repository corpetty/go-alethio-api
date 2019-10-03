package alethio

import "context"

type EtherTransfersService service

// EtherTransfers - ether transfer array structure for the /block/{blockhash}/etherTransfers API endpoint
type EtherTransfers struct {
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

// Get populates ether transfers from a url
func (s *EtherTransfersService) Get(ctx context.Context, link string) (*EtherTransfers, error) {
	req, err := s.client.NewURLRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	etherTransfers := new(EtherTransfers)
	_, err = s.client.Do(ctx, req, &etherTransfers)
	return etherTransfers, err
}

// GetNext will return the next page of ether transfers for a given set of ether transfers
func (s *EtherTransfersService) GetNext(ctx context.Context, et *EtherTransfers) (*EtherTransfers, error) {
	return s.Get(ctx, et.Links.Next)
}

// GetPrev will return the previous page of ether transfers for a given set of ether transfers
func (s *EtherTransfersService) GetPrev(ctx context.Context, et *EtherTransfers) (*EtherTransfers, error) {
	return s.Get(ctx, et.Links.Prev)
}
