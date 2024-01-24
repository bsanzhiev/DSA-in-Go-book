package missingnumber

// Hash-Table, using Hash-Table, we can keep track of the
// elements we have already seen and we can find the missing
// element in just one scan.
// Hash-Table insert and find take constant time O(1) so the total Time
// Complexity of the algorithm is O(n) time. Space Complexity is also O(n)
// 1. Define size of data
// 2. Define map
// 3. Populate map from data
// 4. For loop map
func HashFind(data []int) (int, bool) {
	size := len(data)
	dataMap := make(map[int]bool)

	for _, num := range data {
		dataMap[num] = true
	}

	for i := 1; i <= size+1; i++ {
		if _, found := dataMap[i]; !found {
			return i, true
		}
	}
	return 0, false
}
