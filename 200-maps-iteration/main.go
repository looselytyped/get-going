package main

import "fmt"

func main() {
	// define a map
	codesToStatus := map[int]string{}

	codesToStatus[200] = "OK"
	codesToStatus[404] = "Not found"
	codesToStatus[500] = "Internal Server Error"

	fmt.Printf("codesToStatus is %v\n", codesToStatus)

	// use the range operator to iterate
	// gives you (key, value)
	for k, v := range codesToStatus {
		fmt.Printf("At index %v the language is %v\n", k, v)
	}
}
