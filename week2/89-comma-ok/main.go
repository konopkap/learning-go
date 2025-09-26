package main

import "fmt"

func main() {
	myMap := map[string]string{
		"Adam":    "Abc",
		"Barry":   "Cba",
		"Charlie": "Tango",
	}

	search := "Charlie"

	if v, ok := myMap[search]; ok {
		fmt.Println(search, "does exist in the map and has value:", v)
	} else {
		fmt.Println(search, "doesn't exist in the map!!!")
	}
}
