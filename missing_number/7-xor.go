package missingnumber

func XorFind(data []int, dataRange int) (int, bool) {
	size := len(data)
	xorSum := 0

	// Get XOR of all numbers from 1 to dataRange
	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}

	// Loop through the array and get the XOR of elements
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	return xorSum, true
}
