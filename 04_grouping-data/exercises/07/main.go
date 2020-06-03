package main

import "fmt"

func main() {
	p1 := []string{"james", "bond", "shaken not stirred"}
	p2 := []string{"miss", "moneypenny", "Hello, james"}
	x := [][]string{p1, p2}
	fmt.Println(x)
}
