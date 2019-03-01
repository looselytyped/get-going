package main

import "fmt"

func main() {
	languages := [5]string{
		"Go",
		"Pascal",
		"Smalltalk",
		"Oberon",
		"C",
	}

	// NOT AN ASSIGNMENT!!
	// This COPIES the array!
	otherLangs := languages

	otherLangs[1] = "Squeak"
	fmt.Println(languages[1] == "Pascal")
}
