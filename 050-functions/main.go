package main

import (
	"fmt"
	"math"
)

func hypot(x, y int) float64 {
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

func main() {
	x, y := 3, 4
	z := hypot(x, y)
	fmt.Printf("z is %[1]v: %[1]T\n", z)
}
