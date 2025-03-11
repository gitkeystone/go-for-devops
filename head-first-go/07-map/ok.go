package main

import "fmt"

func main() {
	counters := map[string]int{
		"a": 3,
		"b": 0,
	}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	_, ok = counters["c"]
	fmt.Println(ok)
}
