package main

import "fmt"

func main() {
	arr := [5]int{}

	arr[0] = 4
	arr[1] = 3
	arr[2] = 2
	arr[3] = 1
	arr[4] = 0

	for i, v := range arr {
		fmt.Printf("arr[%v] = %v and it's type of %T\n", i, v, v)
	}
}
