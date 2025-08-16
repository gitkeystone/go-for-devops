package cmd

import (
	"fmt"
	"math/rand"
	"time"
)

const CharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

var length int

func genPasswd() {
	rand.Seed(time.Now().UnixNano())

	if length < 8 {
		fmt.Println("Warning: Password length less than 8 is not recommended for security reasons")
		return
	}

	bytes := make([]byte, length)

	// Get one of each required character type
	bytes[0] = getRandomChar("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytes[1] = getRandomChar("abcdefghijklmnopqrstuvwxyz")
	bytes[2] = getRandomChar("0123456789")
	bytes[3] = '_' // Fixed position for underscore for simplicity

	for i := range bytes {
		switch i {
		case 0, 1, 2, 3:
			continue
		default:
			bytes[i] = CharSet[rand.Intn(len(CharSet))]
		}
	}

	fmt.Println(string(bytes))
}

func getRandomChar(chars string) byte {
	return chars[rand.Intn(len(chars))]
}
