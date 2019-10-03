package alethio

import (
	"context"
)

// AccountService is an interface for interfacing with the Account
// endpoints of the Alethio API
// see: https://api.aleth.io/v1/docs#tag/Accounts
type AccountService service

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

// GetDetails will return an ethereum account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}/get
func (s *AccountService) GetDetails(ctx context.Context, address string) (*Account, error) {
	req, err := s.client.NewRequest("GET", "/accounts/"+address, nil)
	if err != nil {
		return nil, err
	}
	account := new(Account)
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetContract takes in an address string and returns the contract details
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1contract/get
func (s *AccountService) GetContract(ctx context.Context, address string) (*Contract, error) {
	req, err := s.client.NewRequest("GET", "accounts/"+address+"/contract", nil)
	if err != nil {
		return nil, err
	}
	contract := new(Contract)
	_, err = s.client.Do(ctx, req, &contract)
	return contract, err
}

// GetTransactions will return the Transactions of a given Account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1transactions/get
func (s *AccountService) GetTransactions(ctx context.Context, address string) (*Transactions, error) {
	req, err := s.client.NewRequest("GET", "accounts/"+address+"/transactions", nil)
	if err != nil {
		return nil, err
	}
	transactions := new(Transactions)
	_, err = s.client.Do(ctx, req, &transactions)
	return transactions, err
}

// GetContractMessages will return the ContractMessages of a given Account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1contractMessages/get
func (s *AccountService) GetContractMessages(ctx context.Context, address string) (*ContractMessages, error) {
	req, err := s.client.NewRequest("GET", "accounts/"+address+"/contractMessages", nil)
	if err != nil {
		return nil, err
	}
	contractMessages := new(ContractMessages)
	_, err = s.client.Do(ctx, req, &contractMessages)
	return contractMessages, err
}

// GetEtherTransfers will return all Ether transfers of a given Account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1etherTransfers/get
func (s *AccountService) GetEtherTransfers(ctx context.Context, address string) (*EtherTransfers, error) {
	req, err := s.client.NewRequest("GET", "accounts/"+address+"/etherTransfers", nil)
	if err != nil {
		return nil, err
	}
	etherTransfers := new(EtherTransfers)
	_, err = s.client.Do(ctx, req, &etherTransfers)
	return etherTransfers, err
}

// GetTokenTransfers will return all Token transfers of a given Account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1tokenTransfers/get
func (s *AccountService) GetTokenTransfers(ctx context.Context, address string) (*TokenTransfers, error) {
	req, err := s.client.NewRequest("GET", "accounts/"+address+"/tokenTransfers", nil)
	if err != nil {
		return nil, err
	}
	tokenTransfers := new(TokenTransfers)
	_, err = s.client.Do(ctx, req, &tokenTransfers)
	return tokenTransfers, err
}
