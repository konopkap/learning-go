package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Run %d\nx is 3\n---------\n", i)
		}
	}
}
