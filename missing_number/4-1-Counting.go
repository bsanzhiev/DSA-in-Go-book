package missingnumber

func FindMissingNumber4(data []int) (int, bool) {
	size := len(data)
	count := make([]int, size)

	for _, num := range data {
		if num <= size {
			count[num-1] = 1
		}
	}
	for i, c := range count {
		if c == 0 {
			return i + 1, true
		}
	}
	return 0, false
}
