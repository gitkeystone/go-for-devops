package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 遇事不决，先写注释

func sayHello(w http.ResponseWriter, _ *http.Request) {
	//  1. 定义模板：hello.tmpl
	//	2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err：%v\n", err)
		return
	}

	//	3. 渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(`:9000`, nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
