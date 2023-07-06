// Двоичная запись
// Дано натуральное число N. Выведите его представление в двоичном виде.
// Входные данные
// Задано единственное число N
// Выходные данные
// Необходимо вывести требуемое представление числа N.

// Sample Input:
// 12
// Sample Output:
// 1100

package main

import (
	"fmt"
)

func GetBinaryView(num int) {
	if num == 0 {
		return
	}

	GetBinaryView(num / 2)
	fmt.Print(num % 2)
}

func main() {
	var num int
	fmt.Scan(&num)

	GetBinaryView(num)
}
