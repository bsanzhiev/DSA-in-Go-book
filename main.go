package main

import (
	"fmt"

	"github.com/bsanzhiev/DSA-in-Go-book/missingnumber"
)

func main() {
	data := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}
	intRange := 10
	missingNum, _ := missingnumber.XorFind(data, intRange)
	fmt.Printf("Missing number: %d\n", missingNum)
}
