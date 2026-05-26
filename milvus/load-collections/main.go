package main

import (
	"fmt"
	"os"

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

func main() {
	clusterEndpoint := os.Getenv("CLUSTER_ENDPOINT")
	token := os.Getenv("TOKEN")

	// List collections
	cs, err := ListCollections(clusterEndpoint, token)
	if err != nil {
		fmt.Println("Error occurred while listing collections:", err.Error())
	}

	// Load collections
	collectionStateCount := make(map[string]int, 3)
	for _, collection := range cs.Data {
		// 未加载: LoadStateNotLoad
		// 加载中: LoadStateLoading
		// 已加载: LoadStateLoaded
		ls, err := GetLoadState(clusterEndpoint, token, collection)
		if err != nil {
			fmt.Println("Error occurred while getting load state:", err.Error())
			continue
		}

		collectionStateCount[ls.Data.LoadState]++

		if ls.Data.LoadState == "LoadStateNotLoad" || ls.Data.LoadState == "LoadStateLoading" {
			continue
		}

		// 未加载: LoadStateNotLoad
		err = LoadCollection(clusterEndpoint, token, collection)
		if err != nil {
			fmt.Println("Error occurred while loading collection:", err.Error())
		}
	}
	fmt.Printf("Milvus 集合加载状态统计:\n未加载: %d\n加载中: %d\n已加载: %d\n", collectionStateCount["LoadStateNotLoad"], collectionStateCount["LoadStateLoading"], collectionStateCount["LoadStateLoaded"])
}

func ListCollections(clusterEndpoint, token string) (cs Collections, err error) {
	url := clusterEndpoint + "/v2/vectordb/collections/list"
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization":   "Bearer " + token,
			"Content-Type":    "application/json",
			"Request-Timeout": "10",
		},
		JSON: map[string]string{}, // Body: {}
	}

	resp, err := grequests.Post(url, grequests.FromRequestOptions(ro))
	if err != nil {
		return cs, fmt.Errorf("Unable to make request: %v", resp.Error)
	}
	if resp.Ok != true {
		return cs, fmt.Errorf("Request did not return OK")
	}
	err = resp.JSON(&cs)
	return cs, err
}

func GetLoadState(clusterEndpoint, token, collectionName string) (ls LoadState, err error) {
	url := clusterEndpoint + "/v2/vectordb/collections/get_load_state"
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization":   "Bearer " + token,
			"Content-Type":    "application/json",
			"Request-Timeout": "10",
		},
		JSON: map[string]string{"collectionName": collectionName}, // Body: {"collectionName": "my_collection"}
	}

	resp, err := grequests.Post(url, grequests.FromRequestOptions(ro))
	if err != nil {
		return ls, fmt.Errorf("Unable to make request: %v", resp.Error)
	}
	if resp.Ok != true {
		return ls, fmt.Errorf("Request did not return OK")
	}
	err = resp.JSON(&ls)
	return ls, err
}

func LoadCollection(clusterEndpoint, token, collectionName string) error {
	url := clusterEndpoint + "/v2/vectordb/collections/load"
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization":   "Bearer " + token,
			"Content-Type":    "application/json",
			"Request-Timeout": "10",
		},
		JSON: map[string]string{"collectionName": collectionName}, // Body: {"collectionName": "my_collection"}
	}

	resp, err := grequests.Post(url, grequests.FromRequestOptions(ro))
	if err != nil {
		return fmt.Errorf("Unable to make request: %v", resp.Error)
	}
	if resp.Ok != true {
		return fmt.Errorf("Request did not return OK")
	}
	return nil
}
