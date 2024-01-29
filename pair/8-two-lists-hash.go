package pair

import "fmt"

func TwoListsHashTable(arr1, arr2 []int, value int) bool {
	//size1 := len(arr1)
	//size2 := len(arr2)
	hs := make(map[int]bool)

	//Step 1: Insert all elements of arr2 into the hash table
	for _, elem := range arr2 {
		hs[elem] = true
	}

	// Шаг 2 и 3: Проверка каждого элемента в arr1, существует ли (значение - elem) в хеш-таблице
	for _, elem := range arr1 {
		if _, found := hs[value-elem]; found {
			fmt.Println("Pair found: ", elem, value-elem)
			return true
		}
	}
	fmt.Println("Pair not found")
	return false
}

func TwoListsHashTable2(arr1, arr2 []int, value int) bool {
	size1 := len(arr1)
	size2 := len(arr2)
	hs := make(map[int]bool)

	// add elements of arr2 to the map
	for i := 0; i < size2; i++ {
		hs[arr2[i]] = true
	}

	// Check for each elem in arr1
	for i := 0; i < size1; i++ {
		if _, found := hs[value-arr1[i]]; found {
			fmt.Println("The pair is :", arr1[i], value-arr1[i])
			return true
		}
	}
	fmt.Println("No pair found")
	return false
}
