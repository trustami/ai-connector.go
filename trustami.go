package ai_connector_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// BaseURL is the base URL for the Trustami.AI API
	BaseURL = "https://api.trustami.ai"
)

type FailedResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func makeRequest(url, token string, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
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

	var failed FailedResponse

	err = json.Unmarshal(body, &failed)
	if err == nil && !failed.Status {
		return nil, fmt.Errorf(failed.Error)
	}

	return body, nil
}
