package ai_connector_go

import (
	"encoding/json"
)

type LangDetector struct {
	_token string
}

// Language is a struct that represents the language of a given text
type Language struct {
	Language string `json:"lang"`
}

// NewLangDetector creates a new LangDetector instance with a given token
func NewLangDetector(token string) *LangDetector {
	return &LangDetector{_token: token}
}

// DetectLanguage detects the language of a given text
func (ld *LangDetector) DetectLanguage(text string) (*Language, error) {
	return detectLanguage(ld._token, text)
}

func detectLanguage(token string, text string) (*Language, error) {
	data := map[string]interface{}{
		"text": text,
	}

	body, err := makeRequest(BaseURL+"/lang", token, data)
	if err != nil {
		return nil, err
	}

	var result Language

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
