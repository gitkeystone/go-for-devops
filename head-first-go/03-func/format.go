package main

import "fmt"

func main() {
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23*5)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Values: %+v %+v %+v\n", 1.2, "\t", true)
	fmt.Printf("Values: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")


	fmt.Printf("%#v %#v %#v", "", "\t", "\n")
	fmt.Printf("%v %v %v", "", "\t", "\n")
}
