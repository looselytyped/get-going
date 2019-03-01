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

	// use the range operator to iterate
	// gives you (key, value)
	for i, v := range languages {
		fmt.Printf("At index %v the language is %v\n", i, v)
	}
}
