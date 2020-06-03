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

	m["todd"] = 33

	if v, ok := m["blah"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9, 45}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
