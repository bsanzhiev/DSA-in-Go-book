package main

import "github.com/bsanzhiev/DSA-in-Go-book/searching"

func main() {
	data := []int{10, 25, 3, 42, 15, 26, 77, 38, 49, 70, 23, 23, 2, 2, 3}
	intrange := 100
	searching.CountSearching(data, intrange)
}