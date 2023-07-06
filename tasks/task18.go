// По данному числу N распечатайте все целые
// значения степени двойки, не превосходящие
// N, в порядке возрастания.
// Входные данные
// Вводится натуральное число.
// Выходные данные
// Выведите ответ на задачу.

// Sample Input:
// 50
// Sample Output:
// 1 2 4 8 16 32

package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scan(&number)
	var count float64 = 0

	for i := 0.0; count < number; i++ {
		fmt.Print(math.Pow(2.0, i), " ")
		count += math.Pow(2.0, i)
	}
}
