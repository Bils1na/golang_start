// На вход дается строка, из нее нужно
// сделать другую строку, оставив только
// нечетные символы (считая с нуля)

// Sample Input:
// ihgewlqlkot
// Sample Output:
// hello

package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)

	for index := range word {
		if index%2 == 1 {
			fmt.Print(string(word[index]))
		}
	}
}
