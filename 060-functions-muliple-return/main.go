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
	msg, err := sayHello("Raju")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msg)
}
