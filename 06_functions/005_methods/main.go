package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r reciever) identifier(parameters) (returns(s)) {code}
func (s secretAgent) speak() {
	fmt.Println("The name's,", s.last, s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	sa1.speak()
	sa2.speak()
}
