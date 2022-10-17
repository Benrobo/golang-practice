package main

import "fmt"

func main() {

	// call function here
	myFunc("benaiah")
	list("benaiah", 20)

	// using function as return value
	add := add(22, 22)
	fmt.Println(add)

	fmt.Println(_func(2, 2))

	// multiple return values
	a, b := customFunc(23, "macbook pro")
	fmt.Println(a, b)
	// omitting a value using (_)
	_, y := customFunc(23, "hello")
	fmt.Println(y)

	// multiple arguments
	sum, k := sumAll(1, 2, 3, 4, 5)
	fmt.Println(sum, k)

	// recursive function
	count(1)
	fmt.Println(factoria(5))
}

// create function using the func keyword

func myFunc(name string) {

	fmt.Println(name)
}

// passing multiple parameters
func list(name string, age int) {
	fmt.Println(name, age)
}

// function return values
// If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function:
/*
	func FunctionName(param1 type, param2 type) type {
		// code to be executed
		return output
	}
*/

func add(x int, y int) int {
	value := x + y
	return value
}

// Named Return Values
// In Go, you can name the return values of a function.

// Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name):

func _func(x int, y int) (result int) {
	result = x + y
	return result
}

// Multiple Return Values
// Go functions can also return multiple values.

func customFunc(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// takes in multiple arguments
func sumAll(num ...int) (sum int, keys []int) {
	sum = 0
	for v, k := range num {
		sum += v
		keys = append(keys, k)
	}
	return sum, keys
}

// recursion
//Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

func count(num int) int {
	if num == 11 {
		return num
	}
	fmt.Println(num)
	return count(num + 1)
}

func factoria(num int) (val int) {
	if num > 0 {
		val = num * factoria(num-1)
	} else {
		val = 1
	}
	return
}
