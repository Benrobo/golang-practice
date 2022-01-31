// a package is a way to group functions, and  it's made up of all the files in the same directory
// every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

package main

import "fmt"

func main() {
	print(`
		My Very First Golang Code.
		
	`)
}

func print(txt string) {
	fmt.Print(txt + "\n")
}
