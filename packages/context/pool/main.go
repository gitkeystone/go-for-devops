package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, Id int) {
	for true {
		select {
		case <-ctx.Done():
			// 获取到取消信号
			// 退出
			return
		default:
			// 处理任务
			time.Sleep(time.Second)
			println("worker", Id, "is working")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//	workers 开始工作
	fmt.Println("start working")
	for i := 0; i < 3; i++ {
		go worker(ctx, i)
	}

	//	工作一段时间后，停止工作
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("stop working")

	for i := 0; i < 10; i++ {
		fmt.Println("通道内容：", <-ctx.Done()) // 再 cancel() 执行前，ctx.Done() 是阻塞的，然后，才关闭了 ctx.Done(), 导致后面所有从 ctx.Done() 获取到的内容都是 {}
		time.Sleep(time.Second)
	}
}
