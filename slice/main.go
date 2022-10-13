package main

import "fmt"

func main() {

	/*
		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

		- Using the []datatype{values} format
		- Create a slice from an array
		- Using the make() function

		slice_name := []datatype{values}
	*/

	ages := []int{}

	fmt.Println(ages)

	/*
		The code above declares an empty slice of 0 length and 0 capacity.

		To initialize the slice during declaration, use this:
	*/
	ages1 := []int{20, 19, 18}

	log(ages1)

	// we can retrieve the length and capacity of a slice using the len() and cap() function

	log(len(ages1))
	log(cap(ages1))

	// creating a slice from an array
	/*
		var myarray = [length]datatype{values} // An array
		myslice := myarray[start:end] // A slice made from the array
	*/

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	_slice := arr1[2:4]

	log(_slice)
	log(cap(_slice))
	log(len(_slice))

	/*
		In the example above _slice is a slice with length 2. It is made from arr1 which is an array with length 6.

		The slice starts from the second element of the array which has value 12. The slice can grow to the end of the array. This means that the capacity of the slice is 4.

		If _slice started from element 0, the slice capacity would be 6.
	*/

	/*
		Append Elements To a Slice
		You can append elements to the end of a slice using the append() function:
	*/

	arr3 := []int{2, 5, 7, 11}

	log(arr3)

	arr4 := append(arr3, 23, 54)

	log(arr4)

	/*
		Append One Slice To Another Slice
		To append all the elements of one slice to another slice, use the append() function:

		Syntax
		slice3 = append(slice1, slice2...)

		Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	*/

	newarr := append(arr3, arr4...)

	log(newarr)

	/*
		Memory Efficiency.

		When using slices, Go loads all the underlying elements into the memory.

		If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

		The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

		copy(dest, src)

		The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	_neededNumbers := numbers[:len(numbers)-10]
	_dest := make([]int, len(_neededNumbers))

	copy(_dest, numbers)

	log(_dest)        // [1 2 3 4 5]
	log(cap(_dest))   // 5
	log(cap(numbers)) // 15

	// as you can see, the capacity of the _dest becomes lesser than the original numbers slice.

}

func log(param ...interface{}) {
	fmt.Println(param[0])
}
