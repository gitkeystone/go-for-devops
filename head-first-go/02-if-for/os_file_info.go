// os_file_info reports file size
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	fileInfo, err := os.Stat("my.txt") // 文件元数据
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size()) // byte
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Sys())
	fmt.Println("------")

	fmt.Println(fileInfo)
	fmt.Printf("%v\n", fileInfo)
	fmt.Printf("%+v\n", fileInfo)
	fmt.Println("------")

	fmt.Println(reflect.TypeOf(fileInfo))
}
