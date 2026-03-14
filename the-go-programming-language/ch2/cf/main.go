package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	cfWithStdin()
}

func cfWithStdin() {
	args := make([]string)

	// 扫描器
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		append(args, input.Text())
	}
}

func cfWithArgs() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
