package main

import "fmt"

func main() {
	slice := []int{11, 32, 36, 84, 25}

	for k, v := range slice {
		fmt.Println("Key: ", k, "| Value:", v)
	}
}
