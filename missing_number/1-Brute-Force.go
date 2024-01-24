package missingnumber

import "fmt"

// Given a list of n-1 elements, which are in the range of 1 to n. There are no
// duplicates in the list. One of the integer is missing. Find the missing element.
func FindMissingNumber(data []int) (int, bool) {
	var found int
	size := len(data)

	for i := 1; i <= size; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if data[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i, true
		}
	}
	fmt.Println("No Number Is Missing")
	return 0, false
}
