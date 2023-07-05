// Количество минимумов
// Найдите количество минимальных элементов в
// последовательности.

// Входные данные
// Вводится натуральное число N, а затем N целых
// чисел последовательности.

// Выходные данные
// Выведите количество минимальных элементов
// последовательности.

// Sample Input:
// 3
// 21
// 11
// 4
// Sample Output:
// 1

package main

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)
	array := make([]int, length, length)

	for idx := range array {
		fmt.Scan(&array[idx])
	}

	var min int = array[0]
	var count int

	for idx := range array {
		if array[idx] < min {
			min = array[idx]
			count = 1
		} else if min == array[idx] {
			count += 1
		}
	}

	fmt.Println(count)
}
