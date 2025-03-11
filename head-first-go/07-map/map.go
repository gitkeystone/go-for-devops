package main

import "fmt"

func main() {
	var ranks map[string]int = make(map[string]int)
	ranks["gold"] = 1
	fmt.Println(ranks)
}
