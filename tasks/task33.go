// Дается строка. Нужно удалить все символы,
// которые встречаются более одного раза и
// вывести получившуюся строку

// Sample Input:
// zaabcbd
// Sample Output:
// zcd

package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scan(&word)

	for i := range word {
		if strings.Count(word, string(word[i])) < 2 {
			fmt.Print(string(word[i]))
		}
	}
}
