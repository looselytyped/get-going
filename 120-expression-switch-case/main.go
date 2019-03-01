package main

import "fmt"

func main() {
	// case can use expressions
	dayIndex := 2
	var msg string
	switch {
	case dayIndex == 1 || dayIndex == 7:
		msg = "Yay! It' the weekend!!"
	default:
		msg = "Blech!!"
	}
	fmt.Println(msg)
}
