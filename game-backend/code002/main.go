package main

import (
	"fmt"
)

func main() {
	var num int
	num = 10
	fmt.Println("变量num的值是：", num)

	str := "这是一个字符串"
	fmt.Println("变量str的值是：", str)

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println("a的值：", a, "b的值：", b, "c的值：", c)

	//var bulletCount int
	bulletCount := 10
	fmt.Println("玩家当前子弹数量：", bulletCount)
}
