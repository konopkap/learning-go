package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:               "John",
		lastName:                "Doe",
		favoriteIceCreamFlavors: []string{"strawberry, vanilla"},
	}
	p2 := person{
		firstName:               "Mark",
		lastName:                "Flap",
		favoriteIceCreamFlavors: []string{"Chocolate, plain, blueberry"},
	}

	fmt.Printf("%v %v's favorite flavours: ", p1.firstName, p1.lastName)
	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Printf("%v,", v)
	}
	fmt.Println()
	fmt.Printf("%v %v's favorite flavours: ", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Printf("%v,", v)
	}
}
