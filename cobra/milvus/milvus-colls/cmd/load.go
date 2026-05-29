/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load Collections",
	Long: `When you load a collection, Milvus loads the index files and the raw data of all fields into memory for 
rapid response to searches and queries. Entities inserted after a collection load are automatically indexed and loaded.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get ENV Variables
		clusterEndpoint := os.Getenv("CLUSTER_ENDPOINT")
		token := os.Getenv("TOKEN")

		url := clusterEndpoint + "/v2/vectordb/collections/list"
		ro := createRequestOptions(token, map[string]string{})

		// Get all collections
		resp, err := createMilvusRequest(url, ro)
		if err != nil {
			log.Fatal(err)
		}
		var collections Collections
		err = resp.JSON(&collections)
		if err != nil {
			log.Fatal(err)
		}

		// Load collections

		collectionList := collections.Data
		counts := len(collectionList)
		bar := progressbar.Default(int64(counts))
		for i := 0; i < counts; i++ {
			// 未加载: LoadStateNotLoad
			// 加载中: LoadStateLoading
			// 已加载: LoadStateLoaded
			bar.Add(1)

			url = clusterEndpoint + "/v2/vectordb/collections/get_load_state"
			collectionName := collectionList[i]
			ro = createRequestOptions(token, map[string]string{"collectionName": collectionName})

			resp, err = createMilvusRequest(url, ro)
			if err != nil {
				log.Println(err)
				continue
			}
			var loadState LoadState
			err = resp.JSON(&loadState)
			if err != nil {
				log.Println(err)
				continue
			}

			// Ignore: LoadStateLoading, LoadStateLoaded
			if loadState.Data.LoadState == "LoadStateLoading" || loadState.Data.LoadState == "LoadStateLoaded" {
				continue
			}

			// LOAD COLLECTION
			// 仅处理未加载: LoadStateNotLoad
			url = clusterEndpoint + "/v2/vectordb/collections/load"
			ro = createRequestOptions(token, map[string]string{"collectionName": collectionName})

			_, err = createMilvusRequest(url, ro)
			if err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
