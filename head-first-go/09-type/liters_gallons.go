package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * .264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * .000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters{
    return Milliliters(g * 3785.41)
}

func main() {
    milk := Gallons(2)
    fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
    fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}
