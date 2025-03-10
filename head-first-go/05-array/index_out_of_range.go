package main

import "fmt"

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	// len
	fmt.Println(len(notes))
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	// for...range
	for index, value := range notes {
		fmt.Println(index, value)
	}
}
