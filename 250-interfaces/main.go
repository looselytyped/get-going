package main

import (
	"fmt"
	"io"
)

// Named represents the interface that allows any thing to have a humane
// identifier
type Named interface {
	name() string
}

// Person is a user-defined type
type Person struct {
	FirstName, LastName string
}

func (p Person) name() string {
	return p.FirstName
}

func (p Person) greet(other Named) string {
	return fmt.Sprintf("Hello %s! My name is %v\n", other.name(), p.FirstName)
}

type ReaderWriter interface {
	io.Reader
	io.Writer
}

func main() {
	raju := Person{"Raju", "Gandhi"}
	brian := Person{"Brian", "Sletten"}
	fmt.Printf(raju.greet(brian))
}
