package main

import "fmt"

func main() {
	c1 := struct {
		class   string
		hp      int
		defense int
		spells  []string
	}{
		class:   "mage",
		hp:      350,
		defense: 100,
		spells: []string{
			"Fire Ball",
			"Blizzard",
			"Blink",
		},
	}
	fmt.Println(c1)
	for _, v := range c1.spells {
		fmt.Println(v)
	}
}
