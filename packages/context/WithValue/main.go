package main

import (
	"context"
	"fmt"
)

type traceIDKey string

const key traceIDKey = "trace_id"

func processRequest(ctx context.Context) {
	id, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("无法获取 trace_id")
		return
	}

	fmt.Printf("处理请求开始， trace_id: %s\n", id)
	logSomething(ctx, "请求处理中...")
	fmt.Printf("处理请求结束，trace_id: %s\n", id)
}

func logSomething(ctx context.Context, message string) {
	id, _ := ctx.Value(key).(string)
	fmt.Printf("日志| trace_id: %s |消息: %s\n", id, message)
}

func main() {
	ctx := context.WithValue(context.Background(), key, "xyz-123-abc")

	processRequest(ctx)
}
