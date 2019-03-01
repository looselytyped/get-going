package main

import (
	"errors"
	"fmt"
	"log"
)

func sayHello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Cannot have name with no length")
	}
	return fmt.Sprintf("Hello %s", name), nil
}

func main() {
	// in-line assignments
	if msg, err := sayHello("Raju"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(msg)
	}
}
