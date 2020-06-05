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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("--------------------")

	}

}
