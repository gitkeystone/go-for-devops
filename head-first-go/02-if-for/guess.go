// guess challenges players to guess a random number.
package main

import (
	"fmt"
	"math/rand"
	// "time"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// seconds := time.Now().Unix()
	// fmt.Println(seconds)
	// rand.Seed(seconds) // 从Go1.20开始，被弃用，默认自动使用更随机的源 crypt/rand

	target := rand.Intn(100) + 1 // Intn(N) -> [0, N)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
