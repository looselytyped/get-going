package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func simple() {
	for i := range [5]int{} {
		fmt.Printf("Simple %v\n", i)
	}
}

func routine() {
	defer wg.Done()
	for i := range [5]int{} {
		fmt.Printf("routine %v\n", i)
	}
}

func main() {
	simple() // blocks

	wg.Add(1)
	go routine() //
	wg.Wait()
}
