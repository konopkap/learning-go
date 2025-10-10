package main

import "fmt"

func main() {
	slice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008"}}
	for i, v1 := range slice {
		for _, v2 := range slice[i] {
			fmt.Println(v1, "-", v2)
		}
	}
}
