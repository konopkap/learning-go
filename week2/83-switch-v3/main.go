package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("-----------\nIteration: 1\n")
		switch x {
		case 0:
			fmt.Println("Value is:", 0)
		case 1:
			fmt.Println("Value is:", 1)
		case 2:
			fmt.Println("Value is:", 2)
		case 3:
			fmt.Println("Value is:", 3)
		case 4:
			fmt.Println("Value is:", 4)
		}
	}
}
