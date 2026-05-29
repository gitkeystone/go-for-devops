/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release Collections",
	Long: `Searches and queries are memory-intensive operations. 
To save the cost, you are advised to release the collections that are currently not in use.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		bar := progressbar.Default(int64(counts), "Releasing")
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

			// Ignore: LoadStateNotLoad
			if loadState.Data.LoadState == "LoadStateNotLoad" {
				continue
			}

			// RELEASE COLLECTION
			// 仅处理已加载: LoadStateNotLoad
			url = clusterEndpoint + "/v2/vectordb/collections/release"
			ro = createRequestOptions(token, map[string]string{"collectionName": collectionName})

			_, err = createMilvusRequest(url, ro)
			if err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// releaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// releaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
