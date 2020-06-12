package main

import (
	"fmt"

	"github.com/fomiller/go-programming-course/13_testing-benchmarking/exercises/02/quote"
	"github.com/fomiller/go-programming-course/13_testing-benchmarking/exercises/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
