package main

import (
	"context"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"log"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

func SayHi(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	return nil, Output{Greeting: "Hi " + input.Name}, nil
}

func main() {
	//	创建服务，并注册一个工具
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "V1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, SayHi)

	//	使用标准输入/输出与客户端通信（Cursor、Claude 等常用方式）
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
