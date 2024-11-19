package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	//	定义模板
	//	解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	//	渲染模板
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}

	m1 := map[string]any{
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}
	//t.Execute(w, "小王子")
	//t.Execute(w, u1)
	//t.Execute(w, m1)
	t.Execute(w, map[string]any{
		"u1": u1,
		"m1": m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(`:9000`, nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
