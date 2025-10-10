package main

import "fmt"

func main() {
	slice := []int{}

	for i := 42; i < 52; i++ {
		slice = append(slice, i)
	}
	slice1 := slice[:5]
	slice2 := slice[5:10]
	slice3 := slice[2:7]
	slice4 := slice[1:6]

	fmt.Println("Slice:", slice)
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
	fmt.Println("Slice3:", slice3)
	fmt.Println("Slice4:", slice4)
}
