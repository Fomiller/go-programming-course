package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", c.info)
}

func main() {
	e1 := customErr{"Need more coffee"}
	foo(e1)
}

func foo(e error) {
	log.Println(e)
}
