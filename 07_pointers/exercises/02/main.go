package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{"Forrest"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Bob"
}
