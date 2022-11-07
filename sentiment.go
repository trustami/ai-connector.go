package ai_connector_go

import (
	"encoding/json"
	"net/url"
)

type SentimentAnalyser struct {
	_token string
}

// Sentiment is a struct that represents the sentiment of a given text
type Sentiment struct {
	Score    float64 `json:"score"`
	Polarity string  `json:"polarity"`
}

// NewSentimentAnalyser creates a new SentimentAnalyser instance with a given token
func NewSentimentAnalyser(token string) *SentimentAnalyser {
	return &SentimentAnalyser{_token: token}
}

// AnalyseSentiment analyses the sentiment of a given text
func (sa *SentimentAnalyser) AnalyseSentiment(text string) (*Sentiment, error) {
	return analyseSentiment(sa._token, text)
}

func analyseSentiment(token string, text string) (*Sentiment, error) {
	data := url.Values{}

	data.Set("text", text)

	body, err := makeRequest(BaseURL+"/sent", token, data)
	if err != nil {
		return nil, err
	}

	var result Sentiment

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
