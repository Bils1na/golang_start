// Напишите функцию, находящую наименьшее
// из четырех введённых в этой же функции чисел.
// Входные данные
// Вводится четыре числа.
// Выходные данные
// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input:
// 4 5 6 7
// Sample Output:
// 4

package main

import (
	"fmt"
)

func minimumFromFour() int {
	var min int
	fmt.Scan(&min)
	var number int
	var i int = 0

	for fmt.Scan(&number); i < 2; fmt.Scan(&number) {
		if number < min {
			min = number
		}
		i++
	}

	return min
}

func main() {
	min := minimumFromFour()

	fmt.Println(min)
}
