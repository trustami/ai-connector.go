package ai_connector_go

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// BaseURL is the base URL for the Trustami.AI API
	BaseURL = "https://api.trustami.ai"
)

func makeRequest(url, token string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
