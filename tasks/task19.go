// Номер числа Фибоначчи
// Дано натуральное число A > 1. Определите,
// каким по счету числом Фибоначчи оно является,
// то есть выведите такое число n, что φn=A.
// Если А не является числом Фибоначчи, выведите число -1.

// Входные данные
// Вводится натуральное число.

// Выходные данные
// Выведите ответ на задачу.

// Sample Input:
// 8
// Sample Output:
// 6

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	var fib_f int = 0
	var fib_s int = 1
	var fib_t int = fib_f + fib_s
	var count int = 2

	for i := 0; fib_t < number; i++ {
		fib_f = fib_s
		fib_s = fib_t
		fib_t = fib_f + fib_s
		count++
	}
	if number != fib_t {
		fmt.Print(-1)
	} else {
		fmt.Println(count)
	}
}
