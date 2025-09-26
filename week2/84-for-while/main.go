package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var isReady bool

	for !isReady {
		isReady = rand.Intn(2) == 1
		if isReady {
			fmt.Println("Service is ready!")
		} else {
			fmt.Println("Stil waiting...")
		}
	}
}
