// // VARIABLES

package main

import "fmt"

// // varaiables
// // in go we have diff ways of declearing variables
// // 1. initial type eg ( var a string = "welcome" )
// // 2. inferred type eg ( var a = "welcome" )
// // or without the var keyword ( x := "welcome" )

func main() {

	var n1 string = "benrobo" // type is string
	var n2 = "john"
	n3 := "doe"

	fmt.Println(n1, n2, n3)

	// If a varaiable is declared without a initial value, the value of the variable would be the default type of that varaible

	// var a string
	// var b int
	// var c bool

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	var x string
	x = "ben"

	print(x)

	// // Multiple declearation
	a, b, c := "a", "b", "c"

	print(a)
	print(b)
	print(c)

	// Multiple varaiable in a block
	// var (
	// 	a = 2
	// 	b = "ben"
	// 	c = true
	// )

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Go varaiable naming rules

	myName_2 := "benrobo"

	fmt.Println(myName_2)

}
