package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"weightconv"
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
		w, err := strconv.ParseFloat(line, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		p := weightconv.Pound(w)
		k := weightconv.Kilogram(w)
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.BToK(p), k, weightconv.KToB(k))
	}

}

func cfWithArgs() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		p := weightconv.Pound(w)
		k := weightconv.Kilogram(w)
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.BToK(p), k, weightconv.KToB(k))

	}
}

func main() {
	cfWithArgs()
}
