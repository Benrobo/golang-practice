package main

import "fmt"

func main() {

	// var AGE int = 34
	// var SIZE int = 56
	// total := AGE * SIZE

	var (
		a int
		b int
	)

	a = 10
	b = 2

	// Addition
	log(a + b)

	// Multiplication
	log(a * b)

	// Division
	log(a / b)

	// Modulus
	log(a % b)

}

func log(param ...interface{}) {
	fmt.Println(param)
}
