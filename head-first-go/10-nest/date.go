package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func main() {
	date := Date{}

	err := date.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(28)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

    date = Date{}
    date.year = 2019
    date.month = 14
    date.day = 50
    fmt.Println(date)

    date = Date{year: 0, month: 0, day: -2}
    fmt.Println(date)
}
