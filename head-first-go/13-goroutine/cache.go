package main

import (
	"fmt"
	"time"
)

func sendLetters(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "a"
	time.Sleep(1 * time.Second)
	ch <- "b"
	time.Sleep(1 * time.Second)
	ch <- "c"
	time.Sleep(1 * time.Second)
	ch <- "d"
}

func main() {
	start := time.Now()

	size := 3
	ch := make(chan string, size)
	go sendLetters(ch)

	time.Sleep(5 * time.Second)
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())

	fmt.Println(time.Since(start))
}
