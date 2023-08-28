package ai_connector_go

import (
	"encoding/json"
)

type KeywordExtractor struct {
	_token string
}

// Keyword is a struct that represents a single keyword (index 0 is the keyword (string), index 1 is the score (float64))
type Keyword []interface{}

// Keywords is a struct that represents keywords
type Keywords struct {
	Keywords []Keyword `json:"keywords"`
}

// NewKeywordExtractor creates a new KeywordExtractor instance with a given token
func NewKeywordExtractor(token string) *KeywordExtractor {
	return &KeywordExtractor{_token: token}
}

// Extract extracts keywords from a given text
func (sa *KeywordExtractor) Extract(text string, start, end, topN int, useMMR bool, diversity float64) (*Keywords, error) {
	return extractKeywords(sa._token, text, start, end, topN, useMMR, diversity)
}

func extractKeywords(token, text string, start, end, topN int, useMMR bool, diversity float64) (*Keywords, error) {
	data := map[string]interface{}{
		"text":      text,
		"start":     start,
		"end":       end,
		"top_n":     topN,
		"use_mmr":   useMMR,
		"diversity": diversity,
	}

	body, err := makeRequest(BaseURL+"/keyw", token, data)
	if err != nil {
		return nil, err
	}

	var result Keywords

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
