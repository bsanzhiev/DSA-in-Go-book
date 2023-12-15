package searching

import (
	"fmt"
	"sort"
)

func GetSortMax(data []int) {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1

	// сортируем
	sort.Ints(data)

	// для каждого элемента
	for i := 1; i < size; i++ {
		// сравниваем с соседним слева
		if data[i] == data[i-1] {
			// если равны увеличиваем счетчик на 1
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		// если текущий счетчик больше то
		// он станет макс счетчиком
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	fmt.Println("Max by sort:", max)
	// return max
}
