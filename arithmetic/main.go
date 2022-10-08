package main

import "fmt"

func main() {

	var AGE int = 34
	var SIZE int = 56
	total := AGE * SIZE

	log(total, "welcome")

}

func log(...param interface{}) {
	fmt.Println(param)
}
