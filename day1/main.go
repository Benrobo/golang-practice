// a package is a way to group functions, and  it's made up of all the files in the same directory
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
