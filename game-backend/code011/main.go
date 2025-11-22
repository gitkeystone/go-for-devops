package main

import (
	"fmt"
	"time"
)

//func task(name string) {
//	for i := 0; i < 3; i++ {
//		fmt.Printf("%s: 执行任务 %d\n", name, i)
//		time.Sleep(time.Second)
//	}
//}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("生产者发送：%d\n", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Printf("消费者接收：%d\n", num)
	}
}

func main() {
	// go task("协程1")
	// go task("协程2")
	//
	// time.Sleep(5 * time.Second)
	// fmt.Println("主协程结束")

	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("主协程结束")
}
