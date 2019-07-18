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
