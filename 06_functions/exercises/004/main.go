package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("hi my name is %v, i am %v years old", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Forrest",
		last:  "Miller",
		age:   26,
	}
	p1.speak()
}
