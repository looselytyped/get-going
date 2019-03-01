package main // import "github.com/looselytyped/get-going"

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	m "github.com/looselytyped/introducing-go/chapter8/math"
)

func main() {
	rnd := rand.Intn(10)
	fmt.Printf("The random value is %v\n", rnd)

	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The status is %v\n", resp.StatusCode)

	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
