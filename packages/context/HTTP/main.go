package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler开始处理请求")
	defer log.Println("Handler处理请求结束")

	ctx := r.Context()
	select {
	case <-time.After(10e9):
		fmt.Fprintln(w, "请求处理完毕！")
		log.Println("请求处理完毕！")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("请求被客户端取消!")
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/slow", slowHandler)
	log.Println("服务器启动，监听端口：8080")
	log.Println("请在浏览器访问 http://localhost:8080/slow, 然后在10秒内关闭或停止加载页面")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
