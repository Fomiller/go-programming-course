package main

import "fmt"

var y = 42
var z = "Shaken not stirred"
var a = `James said, "Shaken not stirred"`

func main() {
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
	fmt.Println(a)

}
