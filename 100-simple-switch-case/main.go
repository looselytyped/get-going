package main

import "fmt"

func main() {
	// can switch on the same "type"

	dayIndex := 2
	day := "Sunday"
	switch dayIndex {
	case 1:
		day = "Sunday"
	case 2:
		day = "Monday"
	case 3:
		day = "Tuesday"
	case 4:
		day = "Wednesday"
	case 5:
		day = "Thursday"
	case 6:
		day = "Friday"
	case 7:
		day = "Saturday"
	}
	fmt.Printf("The day is %v\n", day)
}
