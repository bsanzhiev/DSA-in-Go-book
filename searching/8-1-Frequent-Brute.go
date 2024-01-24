package searching

import "fmt"

func GetFrequentBrute(data []int) (int, int) {
	size := len(data)
	mostFreq := 0
	maxCount := 0

	for i := 0; i < size; i++ {
		count := 0
		for j := 0; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			mostFreq = data[i]
			maxCount = count
		}
	}
	fmt.Printf("Most frequent element is %d\n", mostFreq)
	fmt.Printf("Count of most frequent element is %d\n", maxCount)

	return mostFreq, maxCount
}
