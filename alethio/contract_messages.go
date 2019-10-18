package alethio

import (
	"context"
	"fmt"
)

type ContractMessagesService service

// ContractMessage - a single message for a Contract Message
type ContractMessage struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			CmsgIndex         int    `json:"cmsgIndex"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			GlobalRank        []int  `json:"globalRank"`
			MsgCallDepth      int    `json:"msgCallDepth"`
			MsgError          bool   `json:"msgError"`
			MsgErrorString    string `json:"msgErrorString"`
			MsgGasLimit       string `json:"msgGasLimit"`
			MsgGasUsed        int    `json:"msgGasUsed"`
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
				Outputs []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"outputs"`
				Raw string `json:"raw"`
			} `json:"msgPayload"`
			MsgType    string `json:"msgType"`
			TxGasPrice string `json:"txGasPrice"`
			Value      string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
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
			Originator struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"originator"`
			ParentContractMessage struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"parentContractMessage"`
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

// ContractMessages - structure for the /blocks/{blockHash}/contractMessages endpoint
type ContractMessages struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			CmsgIndex         int    `json:"cmsgIndex"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			GlobalRank        []int  `json:"globalRank"`
			MsgCallDepth      int    `json:"msgCallDepth"`
			MsgError          bool   `json:"msgError"`
			MsgErrorString    string `json:"msgErrorString"`
			MsgGasLimit       string `json:"msgGasLimit"`
			MsgGasUsed        int    `json:"msgGasUsed"`
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
			Value      string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
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
			Originator struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"originator"`
			ParentContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"parentContractMessage"`
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

// GetDetails will return all ContractMessages of a given id
// https://api.aleth.io/v1/docs#tag/Contract-Messages
func (s *ContractMessagesService) GetDetails(ctx context.Context, id string) (ContractMessages, error) {
	req, err := s.client.NewRequest("GET", "contract-messages/"+id, nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessages ContractMessages
		return emptyContractMessages, err
	}
	var contractMessages ContractMessages
	_, err = s.client.Do(ctx, req, &contractMessages)
	return contractMessages, err
}

// GetSender will return the sender Account of a ContractMessage
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1from/get
func (s *ContractMessagesService) GetSender(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/from", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetDestination will return the destination Account of a ContractMessage
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1to/get
func (s *ContractMessagesService) GetDestination(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/to", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetOriginator will return the Account of the Transaction sender of a given ContractMessage
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1originator/get
func (s *ContractMessagesService) GetOriginator(ctx context.Context, id string) (Account, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/originator", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = s.client.Do(ctx, req, &account)
	return account, err
}

// GetTransaction will return the Transaction that executed a given ContractMessage
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1transaction/get
func (s *ContractMessagesService) GetTransaction(ctx context.Context, id string) (Transaction, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/transaction", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransaction Transaction
		return emptyTransaction, err
	}
	var transaction Transaction
	_, err = s.client.Do(ctx, req, &transaction)
	return transaction, err
}

// GetBlock will return the Block that a given ContractMessage is included in
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1to/get
func (s *ContractMessagesService) GetBlock(ctx context.Context, id string) (GetBlock, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/includedInBlock", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = s.client.Do(ctx, req, &block)
	return block, err
}

// GetParent will return the destination Account of a ContractMessage
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1to/get
func (s *ContractMessagesService) GetParent(ctx context.Context, id string) (ContractMessage, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/parentContractMessage", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessage ContractMessage
		return emptyContractMessage, err
	}
	var contractMessage ContractMessage
	_, err = s.client.Do(ctx, req, &contractMessage)
	return contractMessage, err
}

// GetCreatedContracts will return a Contract array of a given ContractMessage id
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1createsContracts/get
func (s *ContractMessagesService) GetCreatedContracts(ctx context.Context, id string) (Contracts, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/createsContracts", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContracts Contracts
		return emptyContracts, err
	}
	var contracts Contracts
	_, err = s.client.Do(ctx, req, &contracts)
	return contracts, err
}

// GetLogEntries will return an LogEntry array of a given ContractMessage id
// https://docs.aleth.io/api#tag/Contract-Messages/paths/~1contract-messages~1{id}~1logEntries/get
func (s *ContractMessagesService) GetLogEntries(ctx context.Context, id string) (LogEntries, error) {
	req, err := s.client.NewRequest("GET", "contract-message"+id+"/logEntries", nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntries LogEntries
		return emptyLogEntries, err
	}
	var logEntries LogEntries
	_, err = s.client.Do(ctx, req, &logEntries)
	return logEntries, err
}
