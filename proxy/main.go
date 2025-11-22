package main

// # 编译为 Linux 64位
// set GOOS=linux
// set GOARCH=amd64
// go build -o auth-proxy main.go

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	ipAddress = "18.1.27.38"
	port      = "31456"
)

func init() {
	if len(os.Args) == 3 {
		ipAddress = os.Args[1]
		port = os.Args[2]
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Proxy server starting on:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 目标服务器地址
	targetURL := "http://" + ipAddress + ":" + port + r.URL.Path

	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}

	// 创建新的请求
	proxyReq := &http.Request{
		Method: r.Method,
		URL:    target,
		Header: r.Header,
		Body:   r.Body,
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 复制响应头
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 设置状态码
	w.WriteHeader(resp.StatusCode)

	// 复制响应体
	io.Copy(w, resp.Body)
}
