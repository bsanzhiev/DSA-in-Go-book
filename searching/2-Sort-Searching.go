package searching

import (
	"fmt"
	"sort"
)

func PrintSorting(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Print("Repeating elements from sorted are: ")
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print("", data[i])
		}
	}
	fmt.Println()
}
