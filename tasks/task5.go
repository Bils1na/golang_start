// Дан массив, состоящий из целых чисел. Нумерация
// элементов начинается с 0. Напишите программу,
// которая выведет элементы массива, индексы которых четны (0, 2, 4...).

// Входные данные
// Сначала задано число N — количество элементов в массиве.
// Далее через пробел записаны N чисел — элементы массива.
// Массив состоит из целых чисел.

// Выходные данные
// Необходимо вывести все элементы массива с чётными индексами.

// Sample Input:
// 6
// 4 5 3 4 2 3
// Sample Output:
// 4 3 2

package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Print("Enter number elements >> ")
	fmt.Scan(&length)
	array := make([]string, length, length)
	var res string

	for idx := range array {
		fmt.Print("Enter element >> ")
		fmt.Scan(&array[idx])
		if idx%2 == 0 {
			res += array[idx] + " "
		}
	}

	fmt.Println(res)
}
