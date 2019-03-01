package main

import "fmt"

// Person is a user-defined type
type Person struct {
	FirstName, LastName string
}

func (p Person) byValue() {
	p.FirstName = "Imma Copy"
}

func (p *Person) byRef() {
	p.FirstName = "Imma the real thing!"
}

func main() {
	raju := Person{"Raju", "Gandhi"}
	raju.byValue()
	fmt.Println("raju.FirstName: ", raju.FirstName)

	ptr := &raju
	ptr.byRef()
	fmt.Println("raju.FirstName: ", raju.FirstName)
}
