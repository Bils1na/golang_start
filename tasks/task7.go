// Дано трехзначное число. Переверните его, а затем выведите.

// Формат входных данных
// На вход дается трехзначное число, не оканчивающееся на ноль.

// Формат выходных данных
// Выведите перевернутое число.

// Sample Input:
// 843
// Sample Output:
// 348

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for i := 0; i < 3; i++ {
		fmt.Print(number % 10)
		number /= 10
	}
}
