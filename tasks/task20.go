// Из натурального числа удалить заданную цифру.
// Входные данные
// Вводятся натуральное число и цифра, которую нужно удалить.
// Выходные данные
// Вывести число без заданных цифр.

// Sample Input:
// 38012732
// 3
// Sample Output:
// 801272

package main

import (
	"fmt"
)

func main() {
	var number int
	var digit int
	var new_number int
	var new_digit int
	var temp int = 1

	fmt.Scan(&number, &digit)

	for i := 0; number != 0; i++ {
		new_digit = number % 10
		if new_digit != digit {
			new_digit *= temp
			new_number += new_digit
			temp *= 10
		}
		number /= 10
	}
	fmt.Print(new_number)
}
