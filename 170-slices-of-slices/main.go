package main

import "fmt"

func main() {
	// define a slice
	languages := []string{
		"Go",
		"Pascal",
		"Smalltalk",
		"Oberon",
		"C",
	}
	fmt.Printf("languages is %#v with length %v\n", languages, len(languages))

	// create a new slice from an existing one
	// these share the same underlying array
	subset := languages[2:4]
	fmt.Printf("subset is %#v with length %v\n", subset, len(subset))

	// change one
	subset[0] = "Java"
	fmt.Println(languages[2] == "Java")
}
