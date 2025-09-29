package main

import "fmt"

func main() {
	slice := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "BittersweetChocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy ChocolatePeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy SunbutterCinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry LemonadeSorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "WolverineTracks (GF)"}

	for i, v := range slice {
		fmt.Printf("Flavor no.%v: %v\n", i, v)
	}
}
