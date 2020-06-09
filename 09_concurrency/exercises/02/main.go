package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi my name is,", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Forrest", "Miller"}

	// this wont work
	// saySomething(p1)

	// this will work
	saySomething(&p1)
}
