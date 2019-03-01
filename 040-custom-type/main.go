package main

import "fmt"

// Name is a user defined type
type Name string

func main() {
	var raju Name

	raju = "Raju"
	// raju = 10 // cannot use 10 (type int) as type Name in assignment
	fmt.Printf("y is %[1]v: %[1]T\n", raju)
}
