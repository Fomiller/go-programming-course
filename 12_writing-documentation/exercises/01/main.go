package main

import (
	"fmt"

	"github.com/fomiller/go-programming-course/12_writing-documentation/exercises/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
}
