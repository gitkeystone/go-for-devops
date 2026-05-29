/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Collections",
	Long:  `This command demonstrates how to obtain the name list of all collections in the currently connected database.`,
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
		collectionStateCount := make(map[string]int, 3)

		collectionNames := collections.Data
		counts := len(collectionNames)
		bar := progressbar.Default(int64(counts), "Listing")
		for i := 0; i < counts; i++ {
			// 未加载: LoadStateNotLoad
			// 加载中: LoadStateLoading
			// 已加载: LoadStateLoaded
			bar.Add(1)

			url = clusterEndpoint + "/v2/vectordb/collections/get_load_state"

			collectionName := collectionNames[i]
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

			switch loadState.Data.LoadState {
			case "LoadStateNotLoad":
				collectionStateCount["未加载"]++
			case "LoadStateLoading":
				collectionStateCount["加载中"]++
			case "LoadStateLoaded":
				collectionStateCount["已加载"]++
			default:
				collectionStateCount["未知"]++
			}
		}

		// SHOW
		fmt.Println("Milvus 状态统计:")
		for k, v := range collectionStateCount {
			fmt.Printf("%s: %d\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
