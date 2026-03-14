package main

import (
	"bufio"
	"fmt"
	"lenconv"
	"os"
	"strconv"
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

func showConv(input []string) {
	for _, element := range input {
		num, err := strconv.ParseFloat(element, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := lenconv.Meter(num)
		i := lenconv.Inch(num)

		fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToI(m), i, lenconv.IToM(i))
	}
}

func cfWithStdin() {
	input, err := scanFromStdin()
	if err != nil {
		panic(err)
	}

	showConv(input)
}

func cfWithArgs() {
	showConv(os.Args[1:])
}

func main() {
	cfWithStdin()
}
