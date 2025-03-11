package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 3,
		"b": 0,
	}
	for key := range m {
		fmt.Println(key)
	}

	s := []int{1, 2, 3}
	for index := range s {
		fmt.Println(index)
	}
}
