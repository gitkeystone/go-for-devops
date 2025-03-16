package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
    flag := os.O_RDONLY
	file, err := os.OpenFile("aardvark.txt", flag, os.FileMode(0600))
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}
