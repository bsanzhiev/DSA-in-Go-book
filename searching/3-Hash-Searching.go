package searching

import "fmt"

func HashSearching(data []int) {
	s := make(map[int]bool)
	size := len(data)
	fmt.Print("Repeating elements from hash are:")

	for i := 0; i < size; i++ {
		if _, found := s[data[i]]; found {
			fmt.Print(" ", data[i])
		} else {
			s[data[i]] = true
		}
	}
	fmt.Println()
}
