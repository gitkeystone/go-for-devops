package main

import (
	"context"
	"fmt"
	"time"
)

func monitor(ctx context.Context, name string) {
	fmt.Printf("【%s】监控启动...\n", name)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("【%s】收到取消信号，监控停止。原因: %s\n", name, ctx.Err())
			return
		default:
			fmt.Printf("【%s】正在监控中...\n", name)
			time.Sleep(1e9)
		}
	}
}
func main() {
	//fmt.Println("haha")
	ctx, cancel := context.WithCancel(context.Background())
	go monitor(ctx, "监控1号")
	time.Sleep(5e9)
	cancel()
	time.Sleep(1e9)
	fmt.Println("主程序：退出。")
}
