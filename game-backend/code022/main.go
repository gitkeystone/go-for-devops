package main

import "fmt"

func main() {
	//type Character struct {
	//	Name string
	//	HP   int
	//}
	//
	//var ptr *Character
	//
	//var player = Character{
	//	Name: "赵云",
	//	HP:   1500,
	//}
	//
	//ptr = &player
	//fmt.Printf("player变量的地址: %p\n", &player)
	//fmt.Printf("ptr指针的值: %p\n", ptr)

	// 赋值都是通过复制的方式给变量赋值
	// 它们是独立的个体，有自己专属的内存空间
	//var a, b int
	//a = 20
	//b = a
	//b = 10
	//fmt.Printf("a = %d\n", a)
	//fmt.Printf("b = %d\n", b)

	var a int
	var b *int
	a = 20
	b = &a
	*b = 10
	fmt.Printf("*b = %d\n", *b)
	fmt.Printf("a = %d\n", a)
}
