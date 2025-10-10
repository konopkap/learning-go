package main

import "fmt"

func main() {
	slice := []int{}

	for i := 42; i < 52; i++ {
		slice = append(slice, i)
	}

	for i, v := range slice {
		fmt.Printf("slice[%v] = %v and it's type of %T\n", i, v, v)
	}
}
