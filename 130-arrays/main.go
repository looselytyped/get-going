package main

import "fmt"

func main() {
	// define an array with initialization
	languages := [5]string{
		"Go",
		"Pascal",
		"Smalltalk",
		"Oberon",
		"C",
	}
	fmt.Printf("languages is %#v with length %v\n", languages, len(languages))

	// zero value array ["", "", "", ""]
	var areasOfUse [3]string
	areasOfUse[2] = "DevOps"
	fmt.Printf("areasOfUse is %#v with length %v\n", areasOfUse, len(areasOfUse))
}
