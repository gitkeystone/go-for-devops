/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "将图片转化base64格式",
	Run: func(cmd *cobra.Command, args []string) {
		mimeType, base64Str, err := imageToBase64(path)
		if err != nil {
			panic(err)
		}

		fmt.Printf("data:%s;base64,%s", mimeType, base64Str)
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	base64Cmd.Flags().StringVarP(&path, "path", "", "", "图片路径")
}

var path string

func imageToBase64(path string) (string, string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	buffer := b[:512]
	mimeType := http.DetectContentType(buffer)

	imageMIMETypes := []string{
		"image/jpeg",
		"image/png",
		"image/webp",
		"image/gif",
		"image/bmp",
		"image/x-icon",
	}

	if !slices.Contains(imageMIMETypes, mimeType) {
		return mimeType, "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(b)

	return mimeType, base64Str, nil
}
