package searching

import (
	"fmt"
	"sort"
)

func GetMajoritySorting(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		fmt.Printf("Majority is %d\n", data[majIndex])
		return data[majIndex], true
	}
	fmt.Println("No majority element")
	return 0, false
}
