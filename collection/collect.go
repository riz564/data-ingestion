package data_collection

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"data-ingestion/model"
)

func GetPosts(apiURL string, timeout time.Duration) ([]model.Post, error) {
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("non-200 response")
	}
	var result []model.Post
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
