package main

import "fmt"

// Question 11
const (
	Read = 1 << iota
	Write
	Execute
)

func main() {
	fmt.Println("Question 11")
	q11()
	fmt.Println("--------------")

	fmt.Println("Question 12")
	q12()
	fmt.Println("--------------")

	fmt.Println("Question 13")
	q13()
	fmt.Println("--------------")

	fmt.Println("Question 17")
	q17(3232345.342, KB)
}

func q11() {
	fmt.Printf("Read: %b\n", Read)
	fmt.Printf("Write: %b\n", Write)
	fmt.Printf("Execute: %b\n", Execute)
	fmt.Printf("Read+Write: %b\n", Read|Write)
	fmt.Printf("Read+Execute: %b\n", Read|Execute)
	fmt.Printf("Write+Execute: %b\n", Write|Execute)
	fmt.Printf("Read+Write+Execute: %b\n", Read|Write|Execute)
}

func q12() {
	x1, x2 := 13, 24
	x1_and_x2 := x1 & x2
	x1_or_x2 := x1 | x2
	// There's no XOR operator in Go

	fmt.Println("OPERATION (in decimal) | BINARY | DECIMAL")
	fmt.Printf("%d AND %d | %b | %d\n", x1, x2, x1_and_x2, x1_and_x2)
	fmt.Printf("%d OR %d | %b | %d\n", x1, x2, x1_or_x2, x1_or_x2)
}

func q13() {
	var start float64 = 3.14159
	convert := int(start)
	backConvert := float64(convert)

	fmt.Printf("Start: %T %f\n", start, start)
	fmt.Printf("Int: %T %d\n", convert, convert)
	fmt.Printf("Float64: %T %f\n", backConvert, backConvert)
}

const (
	MaxMemory  = "8GB"
	RetryLimit = 3
)

var (
	serverCount  = 5
	debugEnabled = true
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

// All sizes are handled in bytes
func q17(size float32, unit int) {
	size = size * float32(unit)
	fmt.Println("Current size: ")
	fmt.Printf("%.2fB = ", changeSize(float64(size), B))
	fmt.Printf("%.2fKB = ", changeSize(float64(size), KB))
	fmt.Printf("%.2fMB = ", changeSize(float64(size), MB))
	fmt.Printf("%.2fGB", changeSize(float64(size), GB))
}

func changeSize(value float64, destinationUnit int) float64 {
	destinationSize := value / float64(destinationUnit)

	return destinationSize
}
