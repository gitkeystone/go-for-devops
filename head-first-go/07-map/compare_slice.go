package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s == nil\n", s == nil)
	fmt.Printf("make([]int,1) == nil: %#v\n", make([]int, 0) == nil)
	fmt.Printf("[]int{} == nil: %#v\n", []int{} == nil)
}
