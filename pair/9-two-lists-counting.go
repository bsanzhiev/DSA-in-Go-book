package pair

// https://github.com/Hemant-Jain-Author/Data-Structures-Algorithms-In-Go/blob/master/Searching/Searching.go#L768

import "fmt"

func TwoListsCounting(arr1, arr2 []int, rangeVal, value int) bool {
	size1 := len(arr1)
	size2 := len(arr2)

	count := make([]int, rangeVal+1)

	for i := 0; i < size2; i++ {
		count[arr2[i]] = 1
	}

	for i := 0; i < size1; i++ {
		if count[value-arr1[i]] != 0 {
			fmt.Println("The pair is: ", arr1[i], value-arr1[i])
			return true
		}
	}
	fmt.Println("No pair found.")
	return false
}
