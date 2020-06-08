package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	x := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	xs := []byte(x)
	fmt.Println(xs)

	var users []user
	err := json.Unmarshal(xs, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the user data in go form \n", users)

}
