package main

import (
	//"errors"
	"fmt"
	//"log"
)

func main() {
	//err := errors.New("heigth can't be negative")
    err := fmt.Errorf("a heigth of %0.2f is invalid", -2.33333)

	// fmt.Println(err.Error()) // 多余
	fmt.Println(err) // fmt 包自动检查值是否带Error方法
	//log.Fatal(err)

}
