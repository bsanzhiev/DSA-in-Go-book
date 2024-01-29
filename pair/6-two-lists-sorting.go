package pair

import (
	"fmt"
	"sort"
)

// Сортировка. Отсортируйте все элементы во втором списке Y.
// Для каждого элемента X вы можете увидеть, есть ли этот элемент в Y, используя двоичный поиск.
// Алгоритмы сортировки занимают O(m.log m), а поиск — O(n.log m).
// Временная сложность алгоритма равна O(n.log m) или O(m.log m),
// а пространственная сложность — O(1).

func TwoListsSorting(arr1 []int, arr2 []int, value int) bool {
	size1 := len(arr1)
	//size2 := len(arr2)
	var ret bool

	sort.Ints(arr2)
	for i := 0; i < size1; i++ {
		if BinarySearch(arr2, value-arr1[i]) {
			fmt.Println("The pair is: ", arr1[i], value-arr1[i])
		}
		ret = true
	}
	if !ret {
		fmt.Println("No pair found")
	}
	return ret
}

// Binary Search Algorithm - Iterative Way

func BinarySearch(data []int, value int) bool {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
