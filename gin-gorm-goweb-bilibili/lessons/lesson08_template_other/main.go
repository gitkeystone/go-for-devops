package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	//渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func xss(w http.ResponseWriter, _ *http.Request) {
	//定义模板
	//解析模板之前，先定义一个函数 safe
	t := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//解析模板
	t, err := t.ParseFiles("xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	//渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://liwenzhou.com'>李文周的博客</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Printf("render template failed, err: %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
