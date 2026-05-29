package cmd

import (
	"fmt"

	"github.com/levigross/grequests"
)

type Collections struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type LoadStateData struct {
	LoadProgress int    `json:"loadProgress"`
	LoadState    string `json:"loadState"`
	Message      string `json:"message"`
}

type LoadState struct {
	Code int           `json:"code"`
	Data LoadStateData `json:"data"`
}

func createRequestOptions(token string, body any) *grequests.RequestOptions {
	return &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization":   "Bearer " + token,
			"Content-Type":    "application/json",
			"Request-Timeout": "10",
		},
		JSON: body,
	}
}

func createMilvusRequest(url string, ro *grequests.RequestOptions) (*grequests.Response, error) {
	resp, err := grequests.Post(url, grequests.FromRequestOptions(ro))
	if err != nil {
		return nil, fmt.Errorf("Unable to make request: %v", resp.Error)
	}
	if resp.Ok != true {
		return nil, fmt.Errorf("Request did not return OK")
	}
	return resp, nil
}
