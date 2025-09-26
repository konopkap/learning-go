package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(251)
	fmt.Println("Variable: x has value of:", x)

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100")
	case x <= 200:
		fmt.Println("Between 101 and 200")
	case x <= 250:
		fmt.Println("Between 201 and 250")
	}
}
