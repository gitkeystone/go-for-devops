package main

import "fmt"

func greet() {
	fmt.Println("Hello, 欢迎来到Code AI世界！")
}

func add(a, b int) int {
	result := a + b
	return result
}

func calculate(num int) (int, int) {
	square := num * num
	cube := num * num * num
	return square, cube
}

func main() {
	//greet()
	//num1 := 5
	//num2 := 3
	//sum := add(num1, num2)
	//fmt.Printf("%d和%d的和是%d\n", num1, num2, sum)
	//
	number := 2
	sq, cb := calculate(2)
	fmt.Printf("%d的平方是%d, 立方是%d\n", number, sq, cb)
}
