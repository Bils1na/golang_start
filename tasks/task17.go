// По данному числу n закончите фразу "На лугу пасется..."
// одним из возможных продолжений: "n коров", "n корова", "n коровы",
// правильно склоняя слово "корова".
// Входные данные
// Дано число n (0<n<100).
// Выходные данные
// Программа должна вывести введенное число n и одно из слов
// (на латинице): korov, korova или korovy, например, 1 korova,
// 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

// Sample Input:
// 10
// Sample Output:
// 10 korov

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	switch {
	case number > 10 && number < 20:
		fmt.Printf("%d korov ", number)
	case number%10 == 1:
		fmt.Printf("%d korova ", number)
	case number%10 > 0 && number%10 < 5:
		fmt.Printf("%d korovy ", number)
	case number%10 < 1 || number%10 >= 5:
		fmt.Printf("%d korov ", number)
	}

}
