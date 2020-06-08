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
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}
	fmt.Println("BEFORE MARSHAL: \n", users)

	// your code goes here
	userJson, err := json.Marshal(users)
	// log err if needed
	if err != nil {
		fmt.Println(err)
	}
	// print JSON
	fmt.Println("AFTER MARSHAL\n", string(userJson))

}
