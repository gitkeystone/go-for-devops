// dup2 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				log.Printf("dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			for _, n := range counts {
				if n > 1 {
					fmt.Printf("%s\n", f.Name())
					break
				}
			}

			counts = make(map[string]int)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	check(input.Err())
}
