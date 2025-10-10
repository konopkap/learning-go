package main

import "fmt"

func main() {
	slice := []int{}

	for i := 42; i < 52; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 52)
	fmt.Println("Slice:", slice)

	slice = append(slice, 53, 54, 55)
	fmt.Println("Slice:", slice)

	y := []int{56, 57, 58, 59, 60}
	slice = append(slice, y...)
	fmt.Println("Slice:", slice)
}
