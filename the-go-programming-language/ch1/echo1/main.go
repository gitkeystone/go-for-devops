// echo1 输出命令行参数
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// 1.1
	//fmt.Println(os.Args[0])

	fmt.Println(time.Since(start))
}
