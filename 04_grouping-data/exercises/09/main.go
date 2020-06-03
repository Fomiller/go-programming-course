package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_James":      {"guns", "cars", "women"},
		"moneypenny_miss": {"poetry", "danger", "james bond"},
		"no_dr":           {"evil", "puppies", "explosions"},
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println("---------------")
	m["miller_forrest"] = []string{"gainz", "women", "food"}
	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
