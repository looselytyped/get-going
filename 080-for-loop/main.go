package main

import (
	"fmt"
)

func main() {
	// i is local to the for-loop
	for i := 0; i < 10; i++ {
		fmt.Printf("Current index is %v\n", i)
	}

	// as a while loop
	i := 0
	for i < 10 {
		i++
	}

	// infinite
	// for {
	// can use break
	// }

}
