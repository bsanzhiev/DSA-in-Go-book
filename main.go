package main

import (
	"github.com/bsanzhiev/DSA-in-Go-book/pair"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	rangInt := 22
	value := 22

	pair.TwoListsCounting(arr1, arr2, rangInt, value)

	//missingNum, _ := missingnumber.XorFind(data, intRange)
	//fmt.Printf("Missing number: %d\n", missingNum)
}
