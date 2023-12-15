package main

import "github.com/bsanzhiev/DSA-in-Go-book/searching"

func main() {
	data := []int{10, 25, 42, 15, 26, 77, 38, 49, 70, 23, 28, 52, 93, 10}
	intrange := 100
	searching.GetRangeMax(data, intrange)
}
