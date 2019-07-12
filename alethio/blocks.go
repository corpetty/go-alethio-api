package alethio

import (
	"fmt"
)

// GetAllBlocks - structure when asking for all blocks
type GetAllBlocks struct {
	Data  []Block `json:"data"`
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

// GetBlock - struture for endpoints that return a single block
type GetBlock struct {
	Data Block `json:"data"`
	Meta struct {
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
	} `json:"meta"`
}

// Block - base structure for a block
// TODO: expand out each struct as separate item
type Block struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		BlockBeneficiaryReward string `json:"blockBeneficiaryReward"`
		BlockCreationTime      int    `json:"blockCreationTime"`
		BlockDifficulty        string `json:"blockDifficulty"`
		BlockGasLimit          int    `json:"blockGasLimit"`
		BlockGasUsed           int    `json:"blockGasUsed"`
		BlockHash              string `json:"blockHash"`
		Canonical              bool   `json:"canonical"`
		HasBeneficiaryAlias    string `json:"hasBeneficiaryAlias"`
		Number                 int    `json:"number"`
	} `json:"attributes"`
	Relationships struct {
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
		HasBeneficiary struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"hasBeneficiary"`
		LogEntries struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"logEntries"`
		ParentBlock struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"parentBlock"`
		TokenTransfers struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"tokenTransfers"`
		Transactions struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"transactions"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// GetAllBlocks will return all blocks (paginated) starting from most recently mined
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks/get
func (c *Client) GetAllBlocks() (GetAllBlocks, error) {
	req, err := c.newRequest("GET", "/blocks/", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlocks GetAllBlocks
		return emptyBlocks, err
	}
	var blocks GetAllBlocks
	_, err = c.do(req, &blocks)
	return blocks, err
}

// GetBlockByHash will return a block given a blockhash
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}/get
func (c *Client) GetBlockByHash(blockHash string) (GetBlock, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash, nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = c.do(req, &block)
	return block, err
}

// GetBlockByNumber will return a block given a block number
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{number}/get
func (c *Client) GetBlockByNumber(blockNumber int) (GetBlock, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/blocks/%d", blockNumber), nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = c.do(req, &block)
	return block, err
}

// GetBlockByLabel will return a block given a specific label
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{label}/get
func (c *Client) GetBlockByLabel(label string) (GetBlock, error) {
	req, err := c.newRequest("GET", "/blocks/"+label, nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = c.do(req, &block)
	return block, err
}

// GetBlockParent will return the parent of the blockhash given as input
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1parentBlock/get
func (c *Client) GetBlockParent(blockHash string) (GetBlock, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/parentBlock", nil)
	if err != nil {
		fmt.Print(err)
		var emptyBlock GetBlock
		return emptyBlock, err
	}
	var block GetBlock
	_, err = c.do(req, &block)
	return block, err
}

func (c *Client) GetBlockBeneficiary(blockHash string) (Account, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/hasBeneficiary", nil)
	if err != nil {
		fmt.Print(err)
		var emptyAccount Account
		return emptyAccount, err
	}
	var account Account
	_, err = c.do(req, &account)
	return account, err
}
