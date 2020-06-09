package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response1 struct {
	Page   int
	Fruits []string
}
type user struct {
	First     string
	Last      string
	Age       int
	FavMovies []string `json:"Peaches"`
}

func main() {
	// MARSHALING TO JSON
	// bolB, _ := json.Marshal(true)
	// fmt.Println(string(bolB))

	// intB, _ := json.Marshal(42)
	// fmt.Println(string(intB))

	// strB, _ := json.Marshal("Forrest Miller")
	// fmt.Println(string(strB))

	// slice1 := []string{"Car", "Plane", "Boat", "Motorcycle"}
	// sliceB, _ := json.Marshal(slice1)
	// fmt.Println(string(sliceB))

	// map1 := map[string]string{"first": "Forrest", "last": "Miller"}
	// mapB, _ := json.Marshal(map1)
	// fmt.Println(string(mapB))

	// res1 := response1{
	// 	Page:   34,
	// 	Fruits: []string{"apples", "oranges", "bananas"}}
	// res1B, _ := json.Marshal(res1)
	// fmt.Println(string(res1B))

	// UNMARSHALING JSON
	data1 := []byte(`{"Page":34,"Fruits":["apples","oranges","bananas"]}`)
	data2, err := ioutil.ReadFile("./example.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data1))
	fmt.Println(string(data2))

	var user1 user
	err = json.Unmarshal(data2, &user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)
	fmt.Println(user1.FavMovies)

	res, _ := http.Get("https://jsonplaceholder.typicode.com/users")
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(body))
}
