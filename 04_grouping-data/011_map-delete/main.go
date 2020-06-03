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

	m["todd"] = 33
	fmt.Println(m)
	delete(m, "james")
	fmt.Println(m)
}
