package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		//fmt.Println("index:", index, "value:", arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(time.Since(start))
}
