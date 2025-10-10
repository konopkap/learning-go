package main

import "fmt"

func main() {
	slice := []int{}

	for i := 42; i < 52; i++ {
		slice = append(slice, i)
	}

	slice = append(slice[:3], slice[6:]...)
	fmt.Println("Slice:", slice)
}
