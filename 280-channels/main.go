package main

import (
	"fmt"
)

func routine(c chan<- int) {
	for i := range [5]int{} {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go routine(c)

	for e := range c {
		fmt.Println(e)
	}
}
