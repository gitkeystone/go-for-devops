package main

import "fmt"

func main() {
	var notes []string
	fmt.Printf("%#v\n", notes)
	notes = make([]string, 7)
	fmt.Printf("%#v\n", notes)

    primes := make([]int, 5, 10)
    primes[0] = 2
    primes[1] = 3
    fmt.Printf("%#v\n", primes)

    fmt.Println(len(notes))
    fmt.Println(cap(notes))
    fmt.Println(len(primes))
    fmt.Println(cap(primes))

    p := new([]int)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%#v\n", *p)

    letters := []string{"a", "b", "c"}
    for i:=0;i<len(letters);i++{
        fmt.Println(letters[i])
    }
    for _, letter := range letters {
        fmt.Println(letter)
    }

    notes = []string{"do", "re", "mi", "fa", "so", "la", "ti"}
    fmt.Println(notes[3], notes[6], notes[0])
    primes = []int{
        2,
        3,
        5,
    }
    fmt.Println(primes[0], primes[1], primes[2])
}
