package main

import "fmt"

func main() {
	var a [1]int
	fmt.Printf("%#v\n", a)

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}
