package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "forrest",
		last:        "Miller",
		favIceCream: []string{"Strawberry", "Vanilla"},
	}
	p2 := person{
		"Bob",
		"Builder",
		[]string{
			"Chocolate",
			"Coffee",
		},
	}

	fmt.Println(p1)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)

	}
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)

	}

}
