package pair

import "fmt"

// Найдите пару в двух списках
// Даны два списка X и Y.
// Найдите пару элементов (xi, yi) таких,
// что xi ∈ X и yi ∈ Y, где xi + yi = value.

func TwoListsBrute(arr1 []int, arr2 []int, value int) bool {
	size1 := len(arr1)
	size2 := len(arr2)
	var ret bool

	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			if arr1[i]+arr2[j] == value {
				fmt.Println("The pair is: ", arr1[i], arr2[j])
				ret = true
				return ret
			}
		}
	}
	if !ret {
		fmt.Println("No pais found")
	}
	return ret
}
