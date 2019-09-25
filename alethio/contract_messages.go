package alethio

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

// GetContractMessages will return all ContractMessages of a given id
// https://api.aleth.io/v1/docs#tag/Contract-Messages
// func (c *Client) GetContractMessages(id string) (ContractMessages, error) {
// 	req, err := c.newRequest("GET", "contract-messages/"+id, nil)
// 	if err != nil {
// 		fmt.Print(err)
// 		var emptyContractMessages ContractMessages
// 		return emptyContractMessages, err
// 	}
// 	var contractMessages ContractMessages
// 	_, err = c.do(req, &contractMessages)
// 	return contractMessages, err
// }
