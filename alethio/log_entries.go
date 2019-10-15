package alethio

import (
	"context"
	"fmt"
)

// LogEntriesService is the service for logEntires API endpoint
type LogEntriesService service

// GetLogEntries is the struct that details all a group of logEntries
type GetLogEntries struct {
	Data struct {
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
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
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
	Meta struct {
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
	} `json:"meta"`
}

// GetDetails returns the LogEntires of a given LogEntry id
// https://docs.aleth.io/api#tag/Log-Entries
func (s *LogEntriesService) GetDetails(ctx context.Context, id string) (GetLogEntries, error) {
	req, err := s.client.NewRequest("GET", "log-entries/"+id, nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntries GetLogEntries
		return emptyLogEntries, err
	}
	var logEntries GetLogEntries
	_, err = s.client.Do(ctx, req, &logEntries)
	return logEntries, err
}
