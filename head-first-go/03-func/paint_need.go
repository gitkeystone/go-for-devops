package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64 = 10.0

// func paintNeed(heigth float64, width float64) {
func paintNeeded(heigth, width float64) (float64, error) {
	if heigth < 0 {
		return 0, fmt.Errorf("a heigth of %0.2f is invalid", heigth)
	} else if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	area := heigth * width
	return area / metersPerLiter, nil
}

func main() {
	var amount, total float64

	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.0, 3.3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	fmt.Printf("Total: %.2f liters\n", total)
}
