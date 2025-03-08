package main

import "fmt"

func main() {
	x := 1
	// for ;x <= 3; x++ {
    for ;; {
        if x == 2 {
            break
        }

		fmt.Println(x)
	}
	fmt.Println("End:", x)
}
