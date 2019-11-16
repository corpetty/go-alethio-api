package alethio

import (
	"context"
	"fmt"
)

// LogEntriesService is the service for logEntires API endpoint
type LogEntriesService service

// LogEntry struct is for a single log entry endpoint API call
type LogEntry struct {
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

// GetLogEntries is the struct that details all a group of logEntries
type LogEntryDetails struct {
	Data LogEntry `json:"data"`
	Meta struct {
		LatestBlock struct {
			Number            int    `json:"number"`
			BlockCreationTime int    `json:"blockCreationTime"`
			BlockHash         string `json:"blockHash"`
		} `json:"latestBlock"`
	} `json:"meta"`
}

type LogEntries struct {
	Data  []LogEntry `json:"data"`
	Links NavLinks   `json:"links"`
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

type LogEntry struct {
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
}

// GetDetails returns the LogEntires of a given LogEntry id
// https://docs.aleth.io/api#tag/Log-Entries
func (s *LogEntriesService) GetDetails(ctx context.Context, id string) (LogEntryDetails, error) {
	req, err := s.client.NewRequest("GET", "log-entries/"+id, nil)
	if err != nil {
		fmt.Print(err)
		var emptyLogEntries LogEntryDetails
		return emptyLogEntries, err
	}
	var logEntries LogEntryDetails
	_, err = s.client.Do(ctx, req, &logEntries)
	return logEntries, err
}

// Get populates log entries from an url
func (s *LogEntriesService) Get(ctx context.Context, link string) (*LogEntries, error) {
	req, err := s.client.NewURLRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	logEntries := new(LogEntries)
	_, err = s.client.Do(ctx, req, &logEntries)
	return logEntries, err
}
