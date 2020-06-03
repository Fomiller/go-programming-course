package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      32,
		"MoneyPenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["blah"])

	v, ok := m["blah"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["blah"]; ok {
		fmt.Println(v)
	}
}
