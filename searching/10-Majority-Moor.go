package searching

import "fmt"

func GetMajorityMoore(data []int) (int, bool) {
	size := len(data)
	majIndex := 0
	count := 1

	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	candidate := data[majIndex]
	count = 0
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
