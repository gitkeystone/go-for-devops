package main

import "fmt"

func mix(num int, flag bool, strings ...string){
    fmt.Println(num, flag, strings)
}

func serveralInts(numbers ...int) {
	fmt.Printf("%#v\n", numbers) // 当作切片来处理
}

func main() {
	serveralInts(1)
	serveralInts(1, 2)
    serveralInts()
}


