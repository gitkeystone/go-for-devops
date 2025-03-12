package main

import (
	"fmt"
)

type Subscriber struct {
        Name   string
        Rate   float64
        Active bool
}

func PrintInfo(s *Subscriber) {
        fmt.Println("Name:", s.Name)
        fmt.Println("Monthly rate", s.Rate)
        fmt.Println("Active?", s.Active)
}

func DefaultSubscriber(name string) *Subscriber {
        var s Subscriber
        s.Name = name
        s.Rate = 5.99
        s.Active = true
        return &s
}

func ApplyDiscount(s *Subscriber) {
        s.Rate = 4.99 // 使用点运算符在struct指针和struct上都可访问字段
}
func main() {
	s := Subscriber{
        Name: "Aman Singh",
    }
	fmt.Printf("%#v\n", s)
}
