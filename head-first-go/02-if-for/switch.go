package main

import "fmt"

func main() {
	v := 99
	switch v {
	default:
		fmt.Println("default")

	case 1:
		fmt.Println(1)
		fallthrough

	case 2, 3:
		fmt.Println(2, 3)
		fallthrough
	case 4:
		fmt.Println(4)
	}

	// type switch
	i := any(1.0)
	switch t := i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Printf("type: %T\n", t)
	}
}
