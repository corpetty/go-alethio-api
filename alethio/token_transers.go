package alethio

import (
	"context"
	"fmt"
)

// TokenTransfersService is the service for all TokenTransfer API calls
type TokenTransfersService service

// TokenTransfer struct contains details of a single token tranfer object
type TokenTransfer struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Decimals          int    `json:"decimals"`
			GlobalRank        []int  `json:"globalRank"`
			Symbol            string `json:"symbol"`
			Value             string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			ContractMessage struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessage"`
			From struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"from"`
			LogEntry struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntry"`
			Originator struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"originator"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
			Token struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"token"`
			Transaction struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"transaction"`
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

// TokenTransfers - structure for token transfers array for /blocks/{blockHash}/tokenTransfers API endpoint
type TokenTransfers struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Decimals          int    `json:"decimals"`
			GlobalRank        []int  `json:"globalRank"`
			Symbol            string `json:"symbol"`
			Value             string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			ContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"contractMessage"`
			From struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"from"`
			LogEntry struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntry"`
			Originator struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"originator"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
			Token struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"token"`
			Transaction struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
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

// GetDetails returns the Token details of a given token contract address
// https://docs.aleth.io/api#tag/Tokens/paths/~1tokens~1{address}/get
func (s *TokenTransfersService) GetDetails(ctx context.Context, id string) (TokenTransfer, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id, nil)
	if err != nil {
		fmt.Print(err)
		var emptyTokenTransfer TokenTransfer
		return emptyTokenTransfer, err
	}
	var tokenTransfer TokenTransfer
	_, err = s.client.Do(ctx, req, &tokenTransfer)
	return tokenTransfer, err
}

// GetToken returns the TokenDetails of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1token/get
func (s *TokenTransfersService) GetToken(ctx context.Context, id string) (TokenDetails, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/token", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTokenDetails TokenDetails
		return emptyTokenDetails, err
	}
	var tokenDetails TokenDetails
	_, err = s.client.Do(ctx, req, &tokenDetails)
	return tokenDetails, err
}

// GetSender returns the sender Account of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1from/get
func (s *TokenTransfersService) GetSender(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/from", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetReceiver returns the destination Account of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1to/get
func (s *TokenTransfersService) GetReceiver(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/to", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetTransaction returns the Transaction of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1transaction/get
func (s *TokenTransfersService) GetTransaction(ctx context.Context, id string) (Transaction, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/transaction", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransaction Transaction
		return emptyTransaction, err
	}
	var transaction Transaction
	_, err = s.client.Do(ctx, req, &transaction)
	return transaction, err
}

// GetContractMessage returns the ContractMessage of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1contractMessage/get
func (s *TokenTransfersService) GetContractMessage(ctx context.Context, id string) (ContractMessage, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/contractMessage", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessage ContractMessage
		return emptyContractMessage, err
	}
	var contractMessage ContractMessage
	_, err = s.client.Do(ctx, req, &contractMessage)
	return contractMessage, err
}

// GetLogEntry returns the LogEntry of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1from/get
func (s *TokenTransfersService) GetLogEntry(ctx context.Context, id string) (LogEntry, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/logEntry", nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntry LogEntry
		return emptyLogEntry, err
	}
	var logEntry LogEntry
	_, err = s.client.Do(ctx, req, &logEntry)
	return logEntry, err
}

// GetOriginator returns the sender Account of a given TokenTransfer id
// https://docs.aleth.io/api#tag/Token-Transfers/paths/~1token-transfers~1{id}~1oroginator/get
func (s *TokenTransfersService) GetOriginator(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "token-transfers/"+id+"/originator", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}
