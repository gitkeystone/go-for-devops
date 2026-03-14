package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func scanFromStdin() ([]string, error) {
	var sli []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sli = append(sli, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return sli, nil
}

func cfWithStdin() {
	input, err := scanFromStdin()
	if err != nil {
		panic(err)
	}

	for _, line := range input {
		t, err := strconv.ParseFloat(line, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
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

func main() {
	cfWithStdin()
}
