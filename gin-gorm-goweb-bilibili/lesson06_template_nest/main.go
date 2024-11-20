package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, _ *http.Request) {
	//  定义一个函数
	//  要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻，又帅气！", nil
	}
	//	定义模板
	t := template.New("f.tmpl") // 创建一个名字是 f 的模板对象;名字一定要跟模板文件名集合中的一个

	//  在解析之前， 告诉模板引擎，我现在多了个自定义的函数 kua
	t.Funcs(template.FuncMap{"kua99": k})

	//	解析模板
	t, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 数据
	name := "小王子"

	//  渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func demo1(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//数据
	name := "小王子"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(`:9000`, nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
