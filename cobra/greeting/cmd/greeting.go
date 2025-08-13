package cmd

import "fmt"

var (
	name = ""
	age  = 0
)

func greeting(name string, age int) {
	fmt.Printf("%s 你好， 今年 %d 岁\n", name, age)
}
