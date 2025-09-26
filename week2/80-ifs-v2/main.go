package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)

	fmt.Println("Value of x: ", x)
	fmt.Println("Value of y: ", y)

	if x < 4 && y < 4 {
		fmt.Println("Both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 and 6")
	} else if y != 5 {
		fmt.Println("y is different than 5")
	} else {
		fmt.Println("None case matched")
	}
}
