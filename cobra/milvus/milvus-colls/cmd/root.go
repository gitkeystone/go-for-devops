/*
Copyright © 2026 Chen Xiaohui
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var clusterEndpoint = os.Getenv("CLUSTER_ENDPOINT")
var token = os.Getenv("TOKEN")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "milvus-colls",
	Short: "A brief description of your application",
	Long: `Loading a collection is the prerequisite to conducting similarity searches and queries in collections. 
This command focuses on the procedures for loading and releasing a collection.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.milvus-colls.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	if clusterEndpoint == "" || token == "" {
		log.Fatalf("\nPlease Set Env: \n\texport CLUSTER_ENDPOINT=\"http://18.1.27.42:19530\"\n\texport TOKEN=\"root:RFA9aYomWKFlzDAP_8lQ\"")
	}
}
