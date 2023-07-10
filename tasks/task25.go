// Напишите функцию sumInt, принимающую
// переменное количество аргументов типа
// int, и возвращающую количество полученных
// функцией аргументов и их сумму. Пакет
// "fmt" уже импортирован, функция и пакет main объявлены.
// Пример вызова вашей функции:
// a, b := sumInt(1, 0)
// fmt.Println(a, b)
// Результат: 2, 1

package main

import (
	"fmt"
)

func sumInt(num ...int) (int, int) {
	var sum int
	var count int

	for _, elem := range num {
		count++
		sum += elem
	}

	return count, sum
}

func main() {
	count, sum := sumInt(1, 2, 3, 4, 5)

	fmt.Println(count, sum)
}
