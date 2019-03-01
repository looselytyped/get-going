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

	// use the inbuilt append
	// re-allocates memory if needed
	languages = append(languages, "Java", "Ruby")
	fmt.Printf("languages is %#v with length %v\n", languages, len(languages))
}
