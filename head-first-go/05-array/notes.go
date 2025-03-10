package main

import "fmt"
import "time"

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[3])
	fmt.Println(notes[6])
	fmt.Println(notes)
	fmt.Printf("%#v\n", notes)

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])
	fmt.Println(primes)
	fmt.Printf("%#v\n", primes)

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])
	fmt.Println(dates)
    fmt.Printf("%#v\n", dates)

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
	fmt.Println(counters)
}
