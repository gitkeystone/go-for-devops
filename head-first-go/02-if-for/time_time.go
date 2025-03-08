package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var now time.Time = time.Now()

	var year int = now.Year() // 方法是与特定类型的值关联的函数

	fmt.Println(reflect.TypeOf(now.Year()))
	fmt.Println(reflect.TypeOf(now.Month())) // int别名
	fmt.Println(reflect.TypeOf(now.Day()))
	fmt.Println(reflect.TypeOf(now.Hour()))
	fmt.Println(reflect.TypeOf(now.Minute()))
	fmt.Println(reflect.TypeOf(now.Second()))
	fmt.Println(reflect.TypeOf(now.Weekday())) // int 别名，Sunday = 0, ...

	var month time.Month = now.Month()
	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Weekday())
}
