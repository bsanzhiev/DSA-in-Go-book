package missingnumber

// Sorting, Sort all the elements in the list and after this in a
// single scan, we can find the duplicates.
// 1. Sort
// 2. Signle scan
func Sorting(data []int) (int, bool) {
	bubbleSort(data)

	for i, num := range data {
		if num != i+1 {
			return i + 1, true
		}
	}
	return len(data) + 1, true
}

func bubbleSort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
