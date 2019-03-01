package main

import "fmt"

func main() {
	// define a map
	codesToStatus := map[int]string{}

	codesToStatus[200] = "OK"
	codesToStatus[404] = "Not found"
	codesToStatus[500] = "Internal Server Error"

	fmt.Printf("codesToStatus is %v\n", codesToStatus)

	// use the inbuilt make function — can supply an optional capacity
	foods := make(map[string]string)
	foods["eggs"] = "protein"
	foods["banana"] = "fruits"
	foods["kale"] = "vegetable"
	foods["soy"] = "protein"

	fmt.Printf("foods is %v\n", foods)
}
