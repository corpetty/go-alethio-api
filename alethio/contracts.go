package alethio

import (
	"context"
	"fmt"
)

type ContractsService service

// Contract is the Contract struct
type Contract struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Address            string   `json:"address"`
			Balance            string   `json:"balance"`
			Bytecode           string   `json:"bytecode"`
			ConstructorArgs    []string `json:"constructorArgs"`
			CreatedAtTimestamp int      `json:"createdAtTimestamp"`
		} `json:"attributes"`
		Relationships struct {
			Account struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"account"`
			ContractMessages struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessages"`
			CreatedAtBlock struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createdAtBlock"`
			CreatedAtContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"createdAtContractMessage"`
			CreatedAtTransaction struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createdAtTransaction"`
			LogEntries struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntries"`
			Token struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"token"`
			Transactions struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"transactions"`
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

// Contracts - array of Contract objects
type Contracts struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Address            string   `json:"address"`
			Balance            string   `json:"balance"`
			Bytecode           string   `json:"bytecode"`
			ConstructorArgs    []string `json:"constructorArgs"`
			CreatedAtTimestamp int      `json:"createdAtTimestamp"`
		} `json:"attributes"`
		Relationships struct {
			Account struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"account"`
			ContractMessages struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessages"`
			CreatedAtBlock struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createdAtBlock"`
			CreatedAtContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"createdAtContractMessage"`
			CreatedAtTransaction struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createdAtTransaction"`
			LogEntries struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntries"`
			Token struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"token"`
			Transactions struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"transactions"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		Count       int `json:"count"`
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
	} `json:"meta"`
}

// GetDetails will return all Contract details of a given address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}/get
func (s *ContractsService) GetDetails(ctx context.Context, address string) (Contract, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address, nil)
	if err != nil {
		fmt.Print(err)
		var emptyContract Contract
		return emptyContract, err
	}
	var contract Contract
	_, err = s.client.Do(ctx, req, &contract)
	return contract, err
}

// GetAccount will return the account that created a given Contract
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1account/get
func (s *ContractsService) GetAccount(ctx context.Context, address string) (Account, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/account", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetContractToken will return Token information based on a token Contract
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1token/get
// currently broken endpoint
// TODO: implement when fixed

// GetCreationBlock will return the creation block of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtBlock/get
func (s *ContractsService) GetCreationBlock(ctx context.Context, address string) (GetBlock, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/createdAtBlock", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = s.client.Do(ctx, req, &block)
	return block, err
}

// GetCreationTransaction will return the creation Transaction of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtTransaction/get
func (s *ContractsService) GetCreationTransaction(ctx context.Context, address string) (Transaction, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/createdAtTransaction", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransaction Transaction
		return emptyTransaction, err
	}
	var transaction Transaction
	_, err = s.client.Do(ctx, req, &transaction)
	return transaction, err
}

// GetCreationMessage will return the creation ContractMessage of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtContractMessage/get
// TODO: endpoing not working, check into this
func (s *ContractsService) GetCreationMessage(ctx context.Context, address string) (ContractMessages, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/createdAtContractMessage", nil)
	if err != nil {
		fmt.Print(err)
		var emptyMessage ContractMessages
		return emptyMessage, err
	}
	var message ContractMessages
	_, err = s.client.Do(ctx, req, &message)
	return message, err
}

// GetTransactions will return (paginated) Transactions of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1transactions/get
func (s *ContractsService) GetTransactions(ctx context.Context, address string) (Transactions, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/transactions", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransactions Transactions
		return emptyTransactions, err
	}
	var transactions Transactions
	_, err = s.client.Do(ctx, req, &transactions)
	return transactions, err
}

// GetMessages will return (paginated) ContractMessages of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1contractMessages/get
func (s *ContractsService) GetMessages(ctx context.Context, address string) (ContractMessages, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/contractMessages", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessages ContractMessages
		return emptyContractMessages, err
	}
	var contractMessages ContractMessages
	_, err = s.client.Do(ctx, req, &contractMessages)
	return contractMessages, err
}

// GetLogEntries will return (paginaged) LogEntries of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1logEntries/get
func (s *ContractsService) GetLogEntries(ctx context.Context, address string) (LogEntries, error) {
	req, err := s.client.NewRequest("GET", "contracts/"+address+"/logEntries", nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntires LogEntries
		return emptyLogEntires, err
	}
	var logEntries LogEntries
	_, err = s.client.Do(ctx, req, &logEntries)
	return logEntries, err
}
