package main

import "fmt"

func main() {
	// define a map
	codesToStatus := map[int]string{}

	codesToStatus[200] = "OK"
	codesToStatus[404] = "Not found"
	codesToStatus[500] = "Internal Server Error"

	// grab a reference
	mapCopy := codesToStatus
	// affect the copy
	delete(mapCopy, 200)

	// affects the original
	fmt.Println(len(codesToStatus) == 2)
}
