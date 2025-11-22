package main

import "fmt"

func main() {
	// 1. 基础
	//for i := 0; i < 5; i++ {
	//	fmt.Println("当前循环次数：", i)
	//}

	// 2.
	//i := 0
	//for i < 5 {
	//	fmt.Println("当前循环次数：", i)
	//	i++
	//}

	// 3.
	//for {
	//	fmt.Println("这是一个无限循环")
	//}

	// 4.
	for {
		fmt.Println("请输入一个数字（输入0退出）：")
		var num int
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		fmt.Printf("你输入的数字是：%d\n", num)
	}
}
