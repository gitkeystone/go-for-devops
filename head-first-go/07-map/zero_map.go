package main

import "fmt"

func main() {
	counters := make(map[string]int)
	//var counters map[string]int // assignment to entry in nil map: 无法给nil映射添加键值对
	// 访问一个没有赋值过的键，返回一个零值（前提是映射已经初始化了）
	counters["a"]++
	counters["b"]++
	counters["c"]++
	fmt.Println(counters)
}
