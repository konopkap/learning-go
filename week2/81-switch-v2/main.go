package e81

import (
	"fmt"
	"math/rand"
)

func Main() {
	x, y := rand.Intn(10), rand.Intn(10)

	fmt.Println("Value of x: ", x)
	fmt.Println("Value of y: ", y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("Both are less than 4")
	case x > 6 && y > 6:
		fmt.Println("Both are greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is between 4 and 6")
	case y != 5:
		fmt.Println("y is different than 5")
	default:
		fmt.Println("None case matched")
	}
}
