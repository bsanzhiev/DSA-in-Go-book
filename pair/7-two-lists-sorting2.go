package pair

import (
	"fmt"
	"sort"
)

//Сортировка. Шаги следующие:
//1. Отсортируйте элементы X и Y в порядке возрастания.
//2. Возьмите сумму наименьшего элемента X и наибольшего элемента Y.
//3. Если сумма равна значению, мы получили нашу пару.
//4. Если сумма меньше значения, взять следующий элемент X
//5. Если сумма больше значения, возьмите предыдущий элемент Y.
//Алгоритмы сортировки занимают O(n.log n) + O(m.log m) для сортировки, а поиск займет время O(n+m).
//Временная сложность алгоритма равна O(n.log n). Пространственная сложность равна O(1).

func TwoListsSorting2(arr1 []int, arr2 []int, value int) bool {
	size1 := len(arr1)
	size2 := len(arr2)

	first := 0
	second := size2 - 1
	curr := 0
	sort.Ints(arr1)
	sort.Ints(arr2)

	for first < size1 && second >= 0 {
		curr = arr1[first] + arr2[second]
		if curr == value {
			fmt.Println("The pair is: ", arr1[first], arr2[second])
			return true
		} else if curr < value {
			first++
		} else {
			second--
		}
	}
	return false
}
