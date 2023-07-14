// Дана строка, содержащая только арабские цифры.
// Найти и вывести наибольшую цифру.
// Входные данные
// Вводится строка ненулевой длины. Известно также,
// что длина строки не превышает 1000 знаков и строка
// содержит только арабские цифры.
// Выходные данные
// Выведите максимальную цифру, которая встречается
// во введенной строке.

// Sample Input:
// 1112221112
// Sample Output:
// 2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	fmt.Scan(&number)
	max, _ := strconv.Atoi(string(number[0]))

	for i := range number {
		k, _ := strconv.Atoi(string(number[i]))

		if max < k {
			max = k
		}
	}

	fmt.Println(max)
}
