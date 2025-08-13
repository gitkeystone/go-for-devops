/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

// my-cli 是go mod init my-cli时指定的模块名，也是一个目录, cmd 是子目录，也是一个子模块
import "my-cli/cmd"

func main() {
	cmd.Execute()
}
