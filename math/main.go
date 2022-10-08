package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.4
	b := 1.6

	// power
	pow := math.Pow(a, 2)

	// sine
	sin := math.Sin(a)

	// cosine
	cos := math.Cos(b)

	log(pow, sin, cos)
}

func log(param ...interface{}) {
	fmt.Println(param)
}
