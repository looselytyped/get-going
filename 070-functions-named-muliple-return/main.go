package main

import (
	"errors"
	"fmt"
	"log"
)

func sayHello(name string) (msg string, err error) {
	msg = ""
	err = nil
	if len(name) == 0 {
		err = errors.New("Cannot have name with no length")
		return
	}
	msg = fmt.Sprintf("Hello %s", name)
	return
}

func main() {
	msg, err := sayHello("Raju")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msg)
}
