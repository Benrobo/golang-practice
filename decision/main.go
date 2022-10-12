package main

import "fmt"

// in golang, we can make decision using if...then and switch..case

func main() {
	var (
		a = 45
		b = 23
	)

	// if...else

	if a > b && a != b {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println(`A is smaller than B`)
	}

	// switch..case

	active := 1

	switch active {
	case 0:
		fmt.Print(`inactive`)
	case 1:
		fmt.Println(`active`)
	default:
		fmt.Print(`neither active nor inactive`)
	}
}
