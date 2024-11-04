package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
	//纯文本
	//_, _ = fmt.Fprintf(writer, "Hello Golang!")
	//带了标签的HTML语言
	//_, _ = fmt.Fprintf(writer, "<h1>Hello Golang!</h1>")
	//_, _ = fmt.Fprintf(writer, "<h1>Hello Golang!</h1><h2>how are you</h2>")
	//从文件中读取
	bs, _ := os.ReadFile("hello.txt")
	_, _ = fmt.Fprintf(writer, string(bs))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
	}

}
