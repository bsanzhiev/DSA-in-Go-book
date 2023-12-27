package searching

import "fmt"

func GetMajorityBrute(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	if maxCount > size/2 {
		fmt.Println("Majority element is", max)
		return max, true
	}
	fmt.Println("No majority element")
	return 0, false
}
