package main

import "fmt"

func main() {
	score := 85
	if score >= 90 {
		fmt.Println("成绩优异！")
	} else if score >= 70 {
		fmt.Println("成绩良好！")
	} else if score >= 60 {
		fmt.Println("成绩及格！")
	} else {
		fmt.Println("成绩不及格，加油！")
	}
}
