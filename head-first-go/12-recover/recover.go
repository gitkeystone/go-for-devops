package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func main() {
	defer calmDown()
    fmt.Println(nil)
	panic("oh no")
}
