// Дана последовательность, состоящая из целых чисел.
// Напишите программу, которая подсчитывает количество
// положительных чисел среди элементов последовательности.

// Входные данные
// Сначала задано число N — количество элементов в последовательности.
// Далее через пробел записаны N чисел — элементы последовательности.
// Последовательность состоит из целых чисел.

// Выходные данные
// Необходимо вывести единственное число - количество
// положительных элементов в последовательности.

// Sample Input:
// 5
// 1 2 3 -1 -4
// Sample Output:
// 3

package main

import "fmt"

func main() {
	var length int
	fmt.Print("Enter number elements >> ")
	fmt.Scan(&length)
	array := make([]int, length, length)
	var sum int

	for idx := range array {
		fmt.Print("Enter element >> ")
		fmt.Scan(&array[idx])
		if array[idx] > 0 {
			sum += 1
		}
	}

	fmt.Printf("Positive numbers = %d", sum)
}
