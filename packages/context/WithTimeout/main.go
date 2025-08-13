package main

import (
	"context"
	"fmt"
	"time"
)

func callAPI(ctx context.Context) {
	fmt.Println("开始调用API...")
	longTimeRunningTask(ctx)
	fmt.Println("API调用完成。")
}

func longTimeRunningTask(ctx context.Context) {
	select {
	case <-time.After(5e9):
		fmt.Println("任务执行完毕！(如果看到此消息，说明未超时)")
	case <-ctx.Done():
		fmt.Printf("任务被中断！原因: %s\n", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3e9)
	defer cancel()
	callAPI(ctx)
}
