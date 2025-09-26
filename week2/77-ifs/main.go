package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)
	fmt.Println("Variable: x has value of:", x)

	if 0 < x && x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x < 200 {
		fmt.Println("Between 101 and 200")
	} else if x < 250 {
		fmt.Println("Between 201 and 250")
	}
}
