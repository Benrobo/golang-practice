package main

import "fmt"

/*

	A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

	While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

	A struct can be useful for grouping data together to create records.

	To declare a structure in Go, use the type and struct keywords:

	type StructName struct {
		member1 datatype;
		member2 datatype;
		member3 datatype;
	}
*/

type Person struct {
	name    string
	age     int
	hobbies []string
	job     string
	salary  int
}

func main() {

	// struct can be used as a variable type when declared.
	var p1 Person
	var p2 Person

	// accessing and modifying struct

	p1.name = "Benaiah"
	p1.age = 20
	p1.hobbies = []string{"coding", "reading"}
	p1.job = "Software Engineer"
	p1.salary = 150000

	p2.name = "Brad"
	p2.age = 40
	p2.hobbies = []string{"coding", "dancing"}
	p2.job = "Software Engineer"
	p2.salary = 150000

	fmt.Println(p1)
	fmt.Println(p2)

	// passing struct as function arguments.

	// Print Pers1 info by calling a function
	printPerson(p1)

	// Print Pers2 info by calling a function
	printPerson(p2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
