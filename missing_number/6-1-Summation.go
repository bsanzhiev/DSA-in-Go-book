package missingnumber

func SummationFind2(data []int) int {
	n := len(data) + 1
	expectedSum := n * (n + 1) / 2

	actualSum := 0
	for _, val := range data {
		actualSum += val
	}

	return expectedSum - actualSum

	//return 0, false
}
