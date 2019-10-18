package alethio

import (
	"context"
	"fmt"
)

type TransactionsService service

// Transaction - structure for an Ethereum Transaction returned by the alethio API
type Transaction struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			FirstSeen         int    `json:"firstSeen"`
			GlobalRank        []int  `json:"globalRank"`
			MsgError          bool   `json:"msgError"`
			MsgErrorString    string `json:"msgErrorString"`
			MsgGasLimit       string `json:"msgGasLimit"`
			MsgPayload        struct {
				FuncName       string        `json:"funcName"`
				FuncSelector   string        `json:"funcSelector"`
				FuncSignature  string        `json:"funcSignature"`
				FuncDefinition string        `json:"funcDefinition"`
				Inputs         []interface{} `json:"inputs"`
				Outputs        []interface{} `json:"outputs"`
				Raw            string        `json:"raw"`
			} `json:"msgPayload"`
			MsgType    string `json:"msgType"`
			TxGasPrice string `json:"txGasPrice"`
			TxGasUsed  int    `json:"txGasUsed"`
			TxHash     string `json:"txHash"`
			TxIndex    int    `json:"txIndex"`
			TxNonce    int    `json:"txNonce"`
			TxType     string `json:"txType"`
			Value      string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			ContractMessages struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessages"`
			CreatesContracts struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createsContracts"`
			From struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"from"`
			IncludedInBlock struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"includedInBlock"`
			LogEntries struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntries"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
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

// Transactions - structure for a single transaction
type Transactions struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			FirstSeen         int    `json:"firstSeen"`
			GlobalRank        []int  `json:"globalRank"`
			MsgError          bool   `json:"msgError"`
			MsgErrorString    string `json:"msgErrorString"`
			MsgGasLimit       string `json:"msgGasLimit"`
			MsgPayload        struct {
				FuncName       string `json:"funcName"`
				FuncSelector   string `json:"funcSelector"`
				FuncSignature  string `json:"funcSignature"`
				FuncDefinition string `json:"funcDefinition"`
				Inputs         []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"inputs"`
				Outputs []interface{} `json:"outputs"`
				Raw     string        `json:"raw"`
			} `json:"msgPayload"`
			MsgType    string `json:"msgType"`
			TxGasPrice string `json:"txGasPrice"`
			TxGasUsed  int    `json:"txGasUsed"`
			TxHash     string `json:"txHash"`
			TxIndex    int    `json:"txIndex"`
			TxNonce    int    `json:"txNonce"`
			TxType     string `json:"txType"`
			Value      string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			ContractMessages struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"contractMessages"`
			CreatesContracts struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"createsContracts"`
			From struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"from"`
			IncludedInBlock struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"includedInBlock"`
			LogEntries struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"logEntries"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
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

// GetDetails returns the Transaction details of a given Transaction hash
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}/get
func (s *TransactionsService) GetDetails(ctx context.Context, transactionHash string) (Transaction, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash, nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransaction Transaction
		return emptyTransaction, err
	}
	var transaction Transaction
	_, err = s.client.Do(ctx, req, &transaction)
	return transaction, err
}

// GetSender returns the sending Account of a given Transaction hash
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}~1from/get
func (s *TransactionsService) GetSender(ctx context.Context, transactionHash string) (Account, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash+"/from", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetDestination returns the destination Account of a given Transaction hash
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}~1to/get
func (s *TransactionsService) GetDestination(ctx context.Context, transactionHash string) (Account, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash+"/to", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetBlock returns the Block that a given Transaction hash was included in
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}~1includedInBlock/get
func (s *TransactionsService) GetBlock(ctx context.Context, transactionHash string) (GetBlock, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash+"/includedInBlock", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = s.client.Do(ctx, req, &block)
	return block, err
}

// GetCreatedContracts returns all Contracts created from a given Transaction hash
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}~1createsContracts/get
func (s *TransactionsService) GetCreatedContracts(ctx context.Context, transactionHash string) (Contracts, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash+"/createsContracts", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContracts Contracts
		return emptyContracts, err
	}
	var contracts Contracts
	_, err = s.client.Do(ctx, req, &contracts)
	return contracts, err
}

// GetContractMessages returns all ContractMessages created from a given Transaction hash
// https://api.aleth.io/v1/docs#tag/Transactions/paths/~1transactions~1{txHash}~1contractMessages/get
func (s *TransactionsService) GetContractMessages(ctx context.Context, transactionHash string) (ContractMessages, error) {
	req, err := s.client.NewRequest("GET", "transactions/"+transactionHash+"/contractMessages", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessages ContractMessages
		return emptyContractMessages, err
	}
	var contractMessages ContractMessages
	_, err = s.client.Do(ctx, req, &contractMessages)
	return contractMessages, err
}
