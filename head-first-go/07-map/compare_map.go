package main

import "fmt"

func main() {
	var xx map[string]int
	fmt.Printf("%#v\n", xx == nil)
	fmt.Printf("map[string]int{} == nil: %#v\n", map[string]int{} == nil)
	fmt.Printf("make(map[string]int) == nil: %#v\n", make(map[string]int) == nil)
	// map 只能跟nil比较
	//fmt.Printf("make(map[string]int) == map[string]int{}: %#v\n", make(map[string]int) == map[string]int{})
}
