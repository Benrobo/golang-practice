package main

import (
	"fmt"
)

func main() {

	for i := 5; i <= 100; i += 5 {
		// fmt.Println(i)
	}

	// skip  an iteration using continue statement
	for i := 5; i <= 100; i += 5 {
		if i == 25 {
			continue
		}
		// fmt.Println(i)
	}

	// break a loop using break statement
	for i := 5; i <= 100; i += 5 {
		if i == 25 {
			break
		}
		// fmt.Println(i)
	}

	// nested loop
	// Here, the "inner loop" will be executed one time for each iteration of the "outer loop":

	adj := []string{"big", "tasty"}
	fruits := []string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			// fmt.Println(adj[i], fruits[j])
		}
	}

	// The Range Keyword
	// The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

	// The range keyword is used like this:
	/*
		for index, value := range array|slice|map {
			// code to be executed for each iteration
		}
	*/
	// the range keyword has access to both the value and index of an array element
	for idx, val := range fruits {
		fmt.Println(idx, fruits[idx], val)
	}

	// Tip: To only show the value or the index, you can omit the other output using an underscore (_).

	// for _, val := range fruits {
	// 	fmt.Println(val)
	// }

	// iteration : while. go doesnt provide the while syntax.
	//We can use for which passes conditional value.

	i := 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}

}
