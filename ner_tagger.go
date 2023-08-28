package ai_connector_go

import (
	"encoding/json"
)

type NERTagger struct {
	_token string
}

// NamedEntities is a struct that represents recognized named entities
type NamedEntities struct {
	Entities []NamedEntity `json:"recognized_named_entites"`
}

// NamedEntity is a struct that represents a single named entity
type NamedEntity struct {
	Text      string `json:"text"`
	Type      string `json:"type"`
	StartChar int    `json:"start_char"`
	EndChar   int    `json:"end_char"`
}

// NewNERTagger creates a new NERTagger instance with a given token
func NewNERTagger(token string) *NERTagger {
	return &NERTagger{_token: token}
}

// RecognizeEntities extracts all Named Entities in a given text
func (sa *NERTagger) RecognizeEntities(text, language string) (*NamedEntities, error) {
	return recognizeEntities(sa._token, text, language)
}

func recognizeEntities(token, text, language string) (*NamedEntities, error) {
	data := map[string]interface{}{
		"text":     text,
		"language": language,
	}

	body, err := makeRequest(BaseURL+"/nert", token, data)
	if err != nil {
		return nil, err
	}

	var result NamedEntities

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
