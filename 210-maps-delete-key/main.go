package main

import "fmt"

func main() {
	// define a map
	codesToStatus := map[int]string{}

	codesToStatus[200] = "OK"
	codesToStatus[404] = "Not found"
	codesToStatus[500] = "Internal Server Error"

	fmt.Printf("codesToStatus is %v\n", codesToStatus)

	delete(codesToStatus, 200)
	fmt.Printf("The new length is %v\n", len(codesToStatus))
}
