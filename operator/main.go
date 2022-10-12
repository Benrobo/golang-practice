package main

import "fmt"

func main() {

	// && and
	// || or
	// ! not
	var (
		a = 34
		b = 20
	)

	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a > b && a != b)
	fmt.Println(a == b || a < b)
}
