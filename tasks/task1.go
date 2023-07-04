// Дано трехзначное число. Найдите сумму его цифр.

// Формат входных данных
// На вход дается трехзначное число.

// Формат выходных данных
// Выведите одно целое число - сумму цифр введенного числа.

// Sample Input:
// 745
// Sample Output:
// 16
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	var sum int

	for number > 0 {
		sum += number % 10
		number /= 10
	}

	fmt.Println(sum)
}
