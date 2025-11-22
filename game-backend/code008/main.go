package main

import "fmt"

func main() {
	//gameMap := [][]int{
	//	{1, 0, 1, 1},
	//	{0, 1, 0, 1},
	//	{1, 0, 1, 0},
	//}
	//
	//for _, row := range gameMap {
	//	for _, cell := range row {
	//		fmt.Printf("%d", cell)
	//	}
	//	fmt.Println()
	//}

	//numbers := []int{1, 2, 3, 4, 5}
	//subSlice := numbers[1:3]
	//fmt.Println(subSlice)

	mySlice := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
	}
	fmt.Println(mySlice)
}
