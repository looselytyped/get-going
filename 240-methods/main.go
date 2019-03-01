package main

import "fmt"

// Person is a user-defined type
type Person struct {
	FirstName, LastName string
}

func (p Person) fullName() string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

func main() {
	raju := Person{"Raju", "Gandhi"}
	fmt.Printf("Raju's fullname is %v\n", raju.fullName())
}
