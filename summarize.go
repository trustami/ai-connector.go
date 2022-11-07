package ai_connector_go

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Summarizer struct {
	_token string
}

// Summary is a struct that represents the summary of a given text
type Summary struct {
	Summary string `json:"summary"`
}

// NewSummarizer creates a new Summarizer instance with a given token
func NewSummarizer(token string) *Summarizer {
	return &Summarizer{_token: token}
}

// SummarizeText summarizes a given text
func (s *Summarizer) SummarizeText(text string, sentenceCount int) (*Summary, error) {
	return summarizeText(s._token, text, sentenceCount)
}

func summarizeText(token string, text string, sentenceCount int) (*Summary, error) {
	data := url.Values{}

	data.Set("text", text)
	data.Set("sentence_count", strconv.Itoa(sentenceCount))

	body, err := makeRequest(BaseURL+"/summ", token, data)
	if err != nil {
		return nil, err
	}

	var result Summary

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
