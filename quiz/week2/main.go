package main

import "fmt"

func main() {
	// separator := "=========================="
	// fmt.Printf("%v\nQuestion 14\n%v\n", separator, separator)
	// q14()
	// fmt.Println(separator)

	// fmt.Printf("%v\nQuestion 15\n%v\n", separator, separator)
	// q15()
	// fmt.Println(separator)

	// fmt.Printf("%v\nQuestion 16\n%v\n", separator, separator)
	// slice := []int{2, 123, 465, 6}
	// fmt.Println("Slice before q16:", slice)
	// q16(slice)
	// fmt.Println("Slice after q16:", slice)
	// fmt.Println(separator)

	servers := []string{"web1", "db1", "web1", "cache1", "db1", "web2"}
	fmt.Println(removeDuplicates(servers))
}

func q14() {
	initCap := 10
	slice := make([]int, 0, initCap)

	for i := 1; i < 16; i++ {
		slice = append(slice, i)
		currentCap := cap(slice)
		if currentCap != initCap {
			fmt.Println("!!! Array reallocated! Capacity has changed from", initCap, "to", currentCap)
			initCap = currentCap
		}
		fmt.Printf("Operation no. %v: Length: %v Capacity: %v\n", i, len(slice), currentCap)
	}
}

func q15() {
	original := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", original)
	original = append(original[:2], original[3:]...)
	fmt.Println("Original without index 2 (30):", original)
}

func q16(s []int) {
	s[0] = 1111
	s[3] = 4444
}

func removeDuplicates(slice []string) []string {
	result := []string{}

	for _, s := range slice {
		found := false
		for _, r := range result {
			if s == r {
				found = true
				break
			}
		}
		if !found {
			result = append(result, s)
		}
	}

	return result
}
