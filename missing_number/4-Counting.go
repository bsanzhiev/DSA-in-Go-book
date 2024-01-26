package missingnumber

// Counting, we know the range of the input so counting will work.
// As we know that, the elements in the list are in the range 0 to n-1.
// We can reserve a list of length n and when we see an element,
// we can increase its count.
// In just one single scan, we know the missing element.
// Counting approach just uses a list so insert and find take
// constant time O(1) so the total Time Complexity of the
// algorithm is O(n) time.
// Space Complexity for creating count list is also O(n)
//
// Вкратце, FindMissingNumber4 проверяет первое недостающее число
//
//	в последовательности от 1 до длины входного массива.
//
// Если число отсутствует, функция возвращает это число и значение true;
// в противном случае возвращается максимальное целое значение и значение false.
// Функция предполагает, что входной массив содержит различные числа в диапазоне от 1 до size.
func CountingFind(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}

	// Создается новый массив count целых чисел длиной size + 1.
	// Этот массив используется для отслеживания наличия чисел в arr.
	count := make([]int, size+1)

	// Массив счетчиков инициализируется значением -1 по всем индексам.
	// Инициализация с -1 указывает, что изначально предполагается, что все числа отсутствуют.
	for i := range count {
		count[i] = -1
	}

	// Для каждого элемента в data соответствующий индекс в массиве count устанавливается равным 1.
	// Элемент data[i] соответствует индексу data[i]-1 в count.
	// Это связано с тем, что числа в data должны находиться в диапазоне от 1 до size,
	// но индексы массива в Go начинаются с 0.
	// Установка count[arr[i]-1] = 1 означает, что число arr[i] присутствует в массиве.
	for i := 0; i < size; i++ {
		count[data[i]-1] = 1
	}

	// Этот цикл проходит через массив count, чтобы найти первый индекс i,
	// значение которого по-прежнему равно -1,
	// что означает, что число i+1 отсутствует в arr.
	for i := 0; i <= size; i++ {
		if count[i] == -1 {
			return i + 1, true
		}
	}
	// Если функция не находит ни одного пропущенного числа
	// (что означает, что в arr присутствуют все числа от 1 до size),
	// она возвращает math.MaxInt32 (максимально возможное целочисленное значение в Go) и false.
	return 0, false
}
