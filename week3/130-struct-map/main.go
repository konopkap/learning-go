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

	m1 := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m1)

	for _, person := range m1 {
		fmt.Printf("%v %v's favorite flavours: ", person.firstName, person.lastName)
		for _, flavor := range person.favoriteIceCreamFlavors {
			fmt.Printf("%v,", flavor)
		}
		fmt.Println()
	}
}
