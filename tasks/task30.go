// На вход подается строка. Нужно определить,
// является ли она палиндромом. Если строка
// является палиндромом - вывести Палиндром
// иначе - вывести Нет. Все входные строчки
// в нижнем регистре.

// Палиндром — буквосочетание, слово или текст,
// одинаково читающееся в обоих направлениях
// (например, "топот", "око", "заказ").

// Sample Input:
// топот
// Sample Output:
// Палиндром

package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)
	runeWord := []rune(word)
	res := ""

	for i := 0; i < len(runeWord); i++ {
		if runeWord[i] == runeWord[len(runeWord)-1-i] {
			res += string(runeWord[i])
		}
	}

	if res == word {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}
