package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func check(err error) {
	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		log.Fatal(err)
	}
}

func main() {
	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("aardvark.txt", flag, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
}
