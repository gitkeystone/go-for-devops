package main

import "fmt"

func greeting(ch chan string) {
	ch <- "hi"
}

func main() {
	ch1 := make(chan string)
	greeting(ch1)
	//fmt.Println(<-ch1)
}
