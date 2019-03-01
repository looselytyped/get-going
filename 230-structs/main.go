package main

import "fmt"

// Person is a user-defined type
type Person struct {
	FirstName, LastName string
}

func main() {
	raju := Person{"Raju", "Gandhi"}
	fmt.Printf("Raju is %#v\n", raju)

	venkat := Person{
		LastName:  "Subramanian",
		FirstName: "Venkat",
	}
	fmt.Printf("Venkat is %#v\n", venkat)
}
