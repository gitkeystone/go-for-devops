/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"

	"github.com/spf13/cobra"
)

// passCmd represents the pass command
var passCmd = &cobra.Command{
	Use:   "pass",
	Short: "密码生成器",
	Long:  `helper pass [-m length]`,
	Run: func(cmd *cobra.Command, args []string) {
		genPasswd()
	},
}

func init() {
	rootCmd.AddCommand(passCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	passCmd.Flags().IntVarP(&length, "chars", "m", 20, "set the character counts")
}

// 自定义逻辑

const CharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

var length int = 8

func genPasswd() {
	rand.Seed(time.Now().UnixNano())

	// 长度不能小于8
	if length < 8 {
		fmt.Println("Warning: Password length less than 8 is not recommended for security reasons")
		os.Exit(1)
	}

	bytes := make([]byte, length)

	// Get one of each required character type
	bytes[0] = getRandomChar("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytes[1] = getRandomChar("abcdefghijklmnopqrstuvwxyz")
	bytes[2] = getRandomChar("0123456789")
	bytes[3] = '_' // Fixed position for underscore for simplicity

	for i := range bytes {
		switch i {
		case 0, 1, 2, 3:
			continue
		default:
			bytes[i] = CharSet[rand.Intn(len(CharSet))]
		}
	}

	// Shuffle to avoid predictable patterns
	for {
		rand.Shuffle(len(bytes), func(i, j int) { bytes[i], bytes[j] = bytes[j], bytes[i] })
		if unicode.IsLetter(rune(bytes[0])) {
			break
		}
	}

	fmt.Println(string(bytes))
}

func getRandomChar(chars string) byte {
	return chars[rand.Intn(len(chars))]
}
