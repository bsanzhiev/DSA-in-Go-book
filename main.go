package main

import (
	"fmt"

	"github.com/bsanzhiev/DSA-in-Go-book/missingnumber"
)

func main() {
	data := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// intrange := 100
	missingNum, found := missingnumber.Sorting(data)
	if found {
		fmt.Printf("Missing number: %d\n", missingNum)
	} else {
		fmt.Println("No missing number")
	}
}
