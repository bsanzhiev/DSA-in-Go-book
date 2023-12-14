// При подсчете этот подход возможен только в том случае, 
// если мы знаем диапазон входных данных. 
// Если мы это знаем, элементы в списке находятся в диапазоне от 0 до n-1.
// Мы можем зарезервировать список длиной n, и когда мы увидим элемент, 
// мы можем увеличить его количество. Всего за одно сканирование мы узнаем дубликаты. 
// Если мы знаем диапазон элементов, то это самый быстрый способ найти дубликаты.

package searching

import "fmt"

func CountSearching(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)

	for i := 0; i < size; i++ {
		count[i] = 0
	}
	fmt.Print("Repeating elements by counting are: ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}
