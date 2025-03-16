package main

import (
	"fmt"
	"strings"
)

func p() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	fmt.Printf("%3d: %08b\n", 0, 0)
	fmt.Printf("%3d: %08b\n", 1, 1)
	fmt.Printf("%3d: %08b\n", 2, 2)
	fmt.Printf("%3d: %08b\n", 3, 3)
	fmt.Printf("%3d: %08b\n", 4, 4)
	fmt.Printf("%3d: %08b\n", 5, 5)
	fmt.Printf("%3d: %08b\n", 6, 6)
	fmt.Printf("%3d: %08b\n", 7, 7)
	fmt.Printf("%3d: %08b\n", 8, 8)
	fmt.Printf("%3d: %08b\n", 32, 32)
	fmt.Printf("%3d: %08b\n", 64, 64)
	fmt.Printf("%3d: %08b\n", 128, 128)

	p()

	fmt.Printf("false && false == %t\n", false && false)
	fmt.Printf("true && false == %t\n", true && false)
	fmt.Printf("true && true == %t\n", true && true)

	p()

	fmt.Printf("%b & %b == %b\n", 0, 0, 0&0)
	fmt.Printf("%b & %b == %b\n", 0, 1, 0&1)
	fmt.Printf("%b & %b == %b\n", 1, 1, 1&1)

	p()

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 1&3)

	p()

	fmt.Printf("%08b\n", 170)
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", 170&15)
}
