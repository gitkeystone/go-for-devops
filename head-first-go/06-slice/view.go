package main

import "fmt"

func main(){
    underlyingArray := [5]string{"a", "b", "c", "d", "e"}
    slice1 := underlyingArray[0:3]
    fmt.Println(slice1)

    i, j := 1, 4
    slice2 := underlyingArray[i:j]
    fmt.Println(slice2)

    slice3 := underlyingArray[2:5]
    fmt.Println(slice3)

    slice4 := underlyingArray[:3]
    fmt.Println(slice4)

    slice5 := underlyingArray[1:]
    fmt.Println(slice5)

    underlyingArray[1] = "X"
    fmt.Println(underlyingArray)
    fmt.Println(slice1)
}
