package pair

import "fmt"

// Two loops

func BrutePair(data []int, value int) bool {
	size := len(data)
	ret := false

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Println("The pair is: ", data[i], ",", data[j])
				ret = true
				return ret
			}
		}
	}
	if !ret {
		fmt.Println("No pair found")
	}
	return ret
}
