package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type userAPIResponse []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/users")
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println(error)
	}

	var APIResponse userAPIResponse
	err := json.Unmarshal(body, &APIResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(APIResponse[0].Address.Geo.Lat)

}
