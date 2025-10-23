package main

import "fmt"

func main() {
	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "One",
		friends: map[string]int{
			"John": 3,
			"Mark": 7,
		},
		favDrinks: []string{"Mango", "Apple"},
	}

	fmt.Println(s1)
}
