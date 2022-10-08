package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter your name: ")
	fmt.Println("Enter DOB: ")

	var (
		username string
		DOB      string
	)

	// accept input from  users using fmt.Scanf() function
	fmt.Scanf("%s", &username)
	fmt.Scanf("%s", &DOB)

	curryear := 2022
	// convert string passed from stdinput to interger
	convertedDob, err := strconv.Atoi(DOB)

	// check if an error occured
	if err != nil {
		fmt.Println("something went wrong")
	}

	diff := curryear - convertedDob
	// convert int to string back.
	ageDiff := strconv.Itoa(diff)

	message := "Welcome comarade 🤓. this tutorial is meant only for those above 18yrs of age. From the program, your name is " + username + " and you are " + ageDiff + "yrs old 😎"

	fmt.Println(message)
}
