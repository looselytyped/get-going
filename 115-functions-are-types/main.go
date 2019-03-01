package main

import "fmt"

func fib() func() int {
	index, a, b := -1, 0, 1
	return func() int {
		switch index++; index {
		case 0:
			return a
		case 1:
			return b
		default:
			a, b = b, a+b
			return a
		}
	}
}
func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
