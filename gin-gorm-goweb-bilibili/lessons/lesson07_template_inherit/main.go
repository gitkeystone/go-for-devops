package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//数据
	msg := "小王子"
	//渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func home(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}

	//数据
	msg := "小王子"

	//渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func index2(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//数据
	name := "小王子"
	//渲染模板
	err = t.ExecuteTemplate(w, "index2.tmpl", name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func home2(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("templates/home2.tmpl", "templates/base.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//数据
	name := "小王子"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
