package alethio

import "fmt"

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

// GetContractDetails will return all Contract details of a given address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}/get
func (c *Client) GetContractDetails(address string) (Contract, error) {
	req, err := c.newRequest("GET", "contracts/"+address, nil)
	if err != nil {
		fmt.Print(err)
		var emptyContract Contract
		return emptyContract, err
	}
	var contract Contract
	_, err = c.do(req, &contract)
	return contract, err
}

// GetContractAccount will return the account that created a given Contract
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1account/get
func (c *Client) GetContractAccount(address string) (Account, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/account", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = c.do(req, &account)
	return account, err
}

// GetContractToken will return Token information based on a token Contract
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1token/get
// currently broken endpoint
// TODO: implement when fixed

// GetContractCreationBlock will return the creation block of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtBlock/get
func (c *Client) GetContractCreationBlock(address string) (GetBlock, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/createdAtBlock", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = c.do(req, &block)
	return block, err
}

// GetContractCreationTransaction will return the creation Transaction of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtTransaction/get
func (c *Client) GetContractCreationTransaction(address string) (Transaction, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/createdAtTransaction", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransaction Transaction
		return emptyTransaction, err
	}
	var transaction Transaction
	_, err = c.do(req, &transaction)
	return transaction, err
}

// GetContractCreationMessage will return the creation ContractMessage of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1createdAtContractMessage/get
// TODO: endpoing not working, check into this
func (c *Client) GetContractCreationMessage(address string) (ContractMessages, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/createdAtContractMessage", nil)
	if err != nil {
		fmt.Print(err)
		var emptyMessage ContractMessages
		return emptyMessage, err
	}
	var message ContractMessages
	_, err = c.do(req, &message)
	return message, err
}

// GetContractTransactions will return (paginated) Transactions of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1transactions/get
func (c *Client) GetContractTransactions(address string) (Transactions, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/transactions", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransactions Transactions
		return emptyTransactions, err
	}
	var transactions Transactions
	_, err = c.do(req, &transactions)
	return transactions, err
}

// GetContractMessages will return (paginated) ContractMessages of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1contractMessages/get
func (c *Client) GetContractMessages(address string) (ContractMessages, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/contractMessages", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessages ContractMessages
		return emptyContractMessages, err
	}
	var contractMessages ContractMessages
	_, err = c.do(req, &contractMessages)
	return contractMessages, err
}

// GetContractLogEntries will return (paginaged) LogEntries of a given Contract address
// https://api.aleth.io/v1/docs#tag/Contracts/paths/~1contracts~1{address}~1logEntries/get
func (c *Client) GetContractLogEntries(address string) (LogEntries, error) {
	req, err := c.newRequest("GET", "contracts/"+address+"/logEntries", nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntires LogEntries
		return emptyLogEntires, err
	}
	var logEntries LogEntries
	_, err = c.do(req, &logEntries)
	return logEntries, err
}
