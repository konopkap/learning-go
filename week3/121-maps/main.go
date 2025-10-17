package main

import "fmt"

func main() {
	myMap := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	for k, val := range myMap {
		fmt.Println(k, ":")
		for i, v := range val {
			fmt.Println("Favorite thing no.", i, "is", v)
		}
	}

	delete(myMap, "no_dr")
	fmt.Println("-----Deleted no_dr------")

	for k, val := range myMap {
		fmt.Println(k, ":")
		for i, v := range val {
			fmt.Println("Favorite thing no.", i, "is", v)
		}
	}
}
