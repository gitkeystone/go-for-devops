package main

import (
	"fmt"
	"os"
)

func p() {
	fmt.Println("-----------------------------------------------")
}

func main() {
	fmt.Println(os.O_RDONLY, os.O_WRONLY, os.O_RDWR)
	fmt.Println(os.O_APPEND, os.O_CREATE, os.O_EXCL, os.O_SYNC, os.O_TRUNC)

	p()

	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)

	p()

	fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE)
	fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
}
