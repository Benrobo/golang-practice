package main

import "fmt"

// A pointer is a programming language object, whose value refers directly to (or “points to”)
// another value stored elsewhere in the computer memory using its address, ref:

func main() {
	x := 10

	fmt.Println(x, &x)

	/*
		You can see x has a value 10 and its memory address 0xC082002238. Memory address
		can be retrieved using & keyword.

		In Go, we can define data type as Pointer using *, for instance var x *int. We can set
		Pointer variable using *x = value. To set memory address on Pointer variable, we can use
		x = memory_address.
	*/

	// declare pointer

	P := &x // assign x memory address to P
	V := *P // V points to the pointer value of P

	V = 45 // reset value ot V to 45

	fmt.Printf("x address :- %v\n", &x)
	fmt.Printf("x value :- %v\n", x)

	fmt.Printf("P address :- %v\n", &P)
	fmt.Printf("P value :- %v\n", *P)

	fmt.Printf("V address :- %v\n", &V)
	fmt.Printf("V value :- %v\n", V)

}
