package main

import "fmt"

func main() {
	// can switch on multiple values
	dayIndex := 2
	var msg string
	switch dayIndex {
	case 1:
		msg = "Yay! It' Sunday!!"
	case 2, 3, 4, 5, 6:
		msg = "Blech!!"
	case 7:
		msg = "Weekend starts!"
	}
	fmt.Println(msg)
}
