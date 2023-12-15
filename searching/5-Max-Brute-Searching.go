package searching

import "fmt"

func GetBruteMax(data []int) {
	// длина слайса
	size := len(data)

	// наиболее встречающийся элемент
	max := data[0]

	// счетчики
	count := 1    // текущий
	maxCount := 1 // маскимальный

	// Во внешнем цикле мы перебираем элементы слайса
	for i := 0; i < size; i++ {
		count = 1
		// внутри считаем входжения элемента
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		// Если число текущего счетчика больше чем предыдущий
		// он становится новым макс значением
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	fmt.Println("Max:", max)
	// return max
}
