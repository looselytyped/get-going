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

	// use the inbuilt make function
	otherLanguages := make([]string, 2)
	otherLanguages[0] = "Java"
	otherLanguages[1] = "Ruby"
	fmt.Printf("otherLanguages is %#v with length %v\n", otherLanguages, len(otherLanguages))
}
