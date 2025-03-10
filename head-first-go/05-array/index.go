package main

import "fmt"

func main() {
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
	fmt.Println(counters)

	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])
	for i := 0; i <= 7; i++ {
		fmt.Println(i, notes[i])
	}
}
