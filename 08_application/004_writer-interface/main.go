package main

import (
	"fmt"
	"io"
	"os"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("Hello")
	fmt.Fprintln(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "HELLO STRING")
}
