package main

import "fmt"

func main() {

	// array are used to store multiple data of same datatypes.
	// syntax :-
	// var arr_name = [length]<data-type>{values}
	// arR_name := []<data-type>{values} ... here length is automatically infered base on  the value.

	var arr1 = [4]int{1, 2, 3, 4}
	// arr2 := []string{"a", "b", "c"}

	// fmt.Println(arr1, arr2)

	// accessing element of an array
	fmt.Println(arr1[0])
	// changing element of an array
	arr1[0] = 45

	fmt.Println(arr1)

	// Array initialization

	/*
		If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

		Tip: The default value for int is 0, and the default value for string is "".
	*/

	arr2 := [5]int{}              //not initialized
	arr3 := [5]int{1, 2}          //partially initialized
	arr4 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr2) // [0 0 0 0 0]
	fmt.Println(arr3) // [1,2,0,0,0]
	fmt.Println(arr4) // [1,2,3,4,5]

	// findind the length of an array using len()

	fmt.Println(len(arr1))
}
