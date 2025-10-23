package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "Audi",
		model:  "A3",
		doors:  3,
		color:  "white",
	}
	v2 := vehicle{
		engine: engine{electric: false},
		make:   "BMW",
		model:  "3",
		doors:  5,
		color:  "black",
	}
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("Electric? -", v1.engine.electric)
	fmt.Println("Electric? -", v2.engine.electric)
}
