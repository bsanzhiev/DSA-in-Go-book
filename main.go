package main

import (
	"github.com/bsanzhiev/DSA-in-Go-book/pair"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 9
	pair.CountingPair(data, value)

	//missingNum, _ := missingnumber.XorFind(data, intRange)
	//fmt.Printf("Missing number: %d\n", missingNum)
}
