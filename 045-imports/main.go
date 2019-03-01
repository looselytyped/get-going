package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	rnd := rand.Intn(10)
	fmt.Printf("The random value is %v\n", rnd)

	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The status is %v\n", resp.StatusCode)

}
