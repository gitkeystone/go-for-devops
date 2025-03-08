package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	// seconds := time.Now().Unix()
	// fmt.Println(seconds)
	// rand.Seed(seconds) // 从Go1.20开始，被弃用，默认自动使用更随机的源 crypt/rand

	target := rand.Intn(100) + 1 // Intn(N) -> [0, N)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)
}
