package alethio

import (
	"fmt"
)

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
	Links struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
	} `json:"links"`
	Meta struct {
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

// GetAccountDetails will return an ethereum account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}/get
func (c *Client) GetAccountDetails(address string) (Account, error) {
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

// GetAccountContract takes in an address string and returns the contract details
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1contract/get
func (c *Client) GetAccountContract(address string) (Contract, error) {
	req, err := c.newRequest("GET", "accounts/"+address+"/contract", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContract Contract
		return emptyContract, err
	}
	var contract Contract
	_, err = c.do(req, &contract)
	return contract, err
}

// GetAccountTransactions will return the Transactions of a given Account
// https://api.aleth.io/v1/docs#tag/Accounts/paths/~1accounts~1{address}~1transactions/get
func (c *Client) GetAccountTransactions(address string) (Transactions, error) {
	req, err := c.newRequest("GET", "accounts/"+address+"/transactions", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransactions Transactions
		return emptyTransactions, err
	}
	var transactions Transactions
	_, err = c.do(req, &transactions)
	return transactions, err
}
