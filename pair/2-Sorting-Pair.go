package pair

import (
	"fmt"
	"sort"
)

func SortingPair(data []int, value int) bool {
	size := len(data)
	sort.Ints(data)
	first := 0
	second := size - 1
	ret := false

	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			fmt.Printf("The pair is %d, %d\n", data[first], data[second])
			ret = true
			break
		}
		if curr < value {
			first++
		} else {
			second--
		}
	}
	if !ret {
		fmt.Println("No pair found")
	}
	return ret
}
