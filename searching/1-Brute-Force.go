// Полный поиск или грубая сила: для каждого элемента в списке найдите,
// есть ли другой элемент с таким же значением.
// Это делается с помощью двух циклов for:
// первого цикла для выбора элемента и
// второго цикла для поиска его повторяющейся записи.

package searching

import "fmt"

func PrintRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are: ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
	}
	fmt.Println()
}
