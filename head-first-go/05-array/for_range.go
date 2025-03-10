package main

import "fmt"

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	// for...range
	for _, value := range notes {
		fmt.Println(value)
	}
}
