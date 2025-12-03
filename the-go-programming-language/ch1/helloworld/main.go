package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	var a, b, _ = test()
	fmt.Println(a, b)

	for i := range []int{1, 2, 3} {
		fmt.Println(i)
	}
}

func test() (int, string, error) {
	return 1, "", nil
}
