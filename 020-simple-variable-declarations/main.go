package main

import (
	"fmt"
)

// zero values
var name string // ""
var xs [5]int   // [0 0 0 0 0]

var scores = [...]int{ // initialize in-line
	10,
	20,
	30,
}

func main() {
	var z = 10 // type inference
	fmt.Printf("z is %[1]v: %[1]T\n", z)

	y := "It's go! Not golang" // short-hand with type inference
	fmt.Printf("y is %[1]v: %[1]T\n", y)
}
