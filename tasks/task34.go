// Ваша задача сделать проверку подходит
// ли пароль вводимый пользователем под
// заданные требования. Длина пароля должна
// быть не менее 5 символов, он должен
// содержать только арабские цифры и буквы
// латинского алфавита. На вход подается
// строка-пароль. Если пароль соответствует
// требованиям - вывести "Ok", иначе вывести "Wrong password"

// Sample Input:
// fdsghdfgjsdDD1
// Sample Output:
// Ok

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)
	var passwordArray []rune = []rune(password)
	var checkVar string = ""

	for i := range passwordArray {
		if unicode.Is(unicode.Latin, passwordArray[i]) || unicode.IsDigit(passwordArray[i]) {
			checkVar += string(passwordArray[i])
		}
	}

	if password == checkVar && len(password) >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
