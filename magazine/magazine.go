package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

type Employee struct {
	Name string
	Salary float64
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}
