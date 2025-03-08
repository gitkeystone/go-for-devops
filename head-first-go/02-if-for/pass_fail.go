// pass_fail reports whether a grade is passing or failing.
package main

import (
	// "reflect"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(reflect.TypeOf(reader))

	input, err := reader.ReadString('\n') // input 包含回车符
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	// fmt.Println(input)

	grade, err := strconv.ParseFloat(input, 64) // err 赋值
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(grade)

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
