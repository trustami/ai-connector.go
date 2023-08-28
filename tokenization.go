package ai_connector_go

import (
	"encoding/json"
)

type Tokenizer struct {
	_token string
}

// TokenizedText is a struct that represents a tokenized text
type TokenizedText struct {
	Sentences string `json:"sentences"`
}

// NewTokenizer creates a new Tokenizer instance with a given token
func NewTokenizer(token string) *Tokenizer {
	return &Tokenizer{_token: token}
}

// Tokenize tokenizes a given text
func (sa *Tokenizer) Tokenize(text string, advanced, spitCamelCase bool) (*TokenizedText, error) {
	return tokenizeText(sa._token, text, advanced, spitCamelCase)
}

func tokenizeText(token, text string, advanced, spitCamelCase bool) (*TokenizedText, error) {
	data := map[string]interface{}{
		"text":            text,
		"advanced":        advanced,
		"spit_camel_case": spitCamelCase,
	}

	body, err := makeRequest(BaseURL+"/tokn", token, data)
	if err != nil {
		return nil, err
	}

	var result TokenizedText

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
