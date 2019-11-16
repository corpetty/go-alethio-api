package alethio

import (
	"context"
	"fmt"
)

type TokensService service

type TokenDetails struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Decimals    int      `json:"decimals"`
			Name        string   `json:"name"`
			Symbol      string   `json:"symbol"`
			TokenTypes  []string `json:"tokenTypes"`
			TotalSupply string   `json:"totalSupply"`
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

// GetDetails returns the Token details of a given token contract address
// https://docs.aleth.io/api#tag/Tokens/paths/~1tokens~1{address}/get
func (s *TokensService) GetDetails(ctx context.Context, address string) (TokenDetails, error) {
	req, err := s.client.NewRequest("GET", "tokens/"+address, nil)
	if err != nil {
		fmt.Print(err)
		var emptyToken TokenDetails
		return emptyToken, err
	}
	var token TokenDetails
	_, err = s.client.Do(ctx, req, &token)
	return token, err
}

// GetContract returns the Token details of a given token contract address
// https://docs.aleth.io/api#tag/Tokens/paths/~1tokens~1{address}~1contract/get
func (s *TokensService) GetContract(ctx context.Context, address string) (TokenTransfers, error) {
	req, err := s.client.NewRequest("GET", "tokens/"+address+"/transers", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransfers TokenTransfers
		return emptyTransfers, err
	}
	var transfers TokenTransfers
	_, err = s.client.Do(ctx, req, &transfers)
	return transfers, err
}
