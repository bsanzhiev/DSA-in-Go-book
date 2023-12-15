package searching

import "fmt"

func GetRangeMax(data []int, dataRange int) {
	size := len(data)
	max := data[0]
	maxCount := 1
	count := make([]int, dataRange)

	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	fmt.Println("Max by ranging:", max)
	// return max
}
