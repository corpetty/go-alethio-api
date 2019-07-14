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

// LogEntries - structure of the /blocks/{blockHash}/logEntries endpoint
type LogEntries struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Cursor       string `json:"cursor"`
			EventDecoded struct {
				Topic0 string `json:"topic0"`
				Event  string `json:"event"`
				Inputs []struct {
					Name    string `json:"name"`
					Type    string `json:"type"`
					Indexed bool   `json:"indexed,omitempty"`
					Value   string `json:"value"`
				} `json:"inputs"`
			} `json:"eventDecoded"`
			EventDecodedError string   `json:"eventDecodedError"`
			GlobalRank        []int    `json:"globalRank"`
			HasLogTopics      []string `json:"hasLogTopics"`
			LogData           string   `json:"logData"`
		} `json:"attributes"`
		Relationships struct {
			Block struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"block"`
			ContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"contractMessage"`
			LoggedBy struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"loggedBy"`
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

// EtherTransfers - ether transfer array structure for the /block/{blockhash}/etherTransfers API endpoint
type EtherTransfers struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BlockCreationTime int    `json:"blockCreationTime"`
			Cursor            string `json:"cursor"`
			Fee               string `json:"fee"`
			GlobalRank        []int  `json:"globalRank"`
			Total             string `json:"total"`
			TransferType      string `json:"transferType"`
			Value             string `json:"value"`
		} `json:"attributes"`
		Relationships struct {
			Block struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"block"`
			ContractMessage struct {
				Data interface{} `json:"data"`
			} `json:"contractMessage"`
			FeeRecipient struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"feeRecipient"`
			From struct {
				Data interface{} `json:"data"`
			} `json:"from"`
			To struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"to"`
			Transaction struct {
				Data interface{} `json:"data"`
			} `json:"transaction"`
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

// GetBlockBeneficiary will return a blocks beneficiary account
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1hasBeneficiary/get
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

// GetBlockTransactions will return the transactions list of a given block
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1transactions/get
func (c *Client) GetBlockTransactions(blockHash string) (Transactions, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/transactions", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTransactions Transactions
		return emptyTransactions, err
	}
	var transactions Transactions
	_, err = c.do(req, &transactions)
	return transactions, err
}

// GetBlockContractMessages will return all contract messages of a given blockHash's block
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1contractMessages/get
func (c *Client) GetBlockContractMessages(blockHash string) (ContractMessages, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/contractMessages", nil)
	if err != nil {
		fmt.Print(err)
		var emptyContractMessages ContractMessages
		return emptyContractMessages, err
	}
	var contractMessages ContractMessages
	_, err = c.do(req, &contractMessages)
	return contractMessages, err
}

// GetBlockLogEntries will return all log entires of a given blockHash's block
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1logEntries/get
func (c *Client) GetBlockLogEntries(blockHash string) (LogEntries, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/logEntries", nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntires LogEntries
		return emptyLogEntires, err
	}
	var logEntries LogEntries
	_, err = c.do(req, &logEntries)
	return logEntries, err
}

// GetBlockEtherTransfers will return all ether transfers of a given blockHash's block
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1etherTransfers/get
func (c *Client) GetBlockEtherTransfers(blockHash string) (EtherTransfers, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/etherTransfers", nil)
	if err != nil {
		fmt.Print(err)
		var emptyEtherTransfers EtherTransfers
		return emptyEtherTransfers, err
	}
	var etherTransfers EtherTransfers
	_, err = c.do(req, &etherTransfers)
	return etherTransfers, err
}

// GetBlockTokenTransfers will return all token transfers of a given blockHash's block
// https://api.aleth.io/v1/docs#tag/Blocks/paths/~1blocks~1{blockHash}~1tokenTransfers/get
func (c *Client) GetBlockTokenTransfers(blockHash string) (TokenTransfers, error) {
	req, err := c.newRequest("GET", "/blocks/"+blockHash+"/tokenTransfers", nil)
	if err != nil {
		fmt.Print(err)
		var emptyTokenTransfers TokenTransfers
		return emptyTokenTransfers, err
	}
	var tokenTransfers TokenTransfers
	_, err = c.do(req, &tokenTransfers)
	return tokenTransfers, err
}

// TODO: impliment the blocks?filter[canonical]=[value] endpoint
