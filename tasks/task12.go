// Количество нулей
// По данным числам, определите количество чисел,
// которые равны нулю.

// Входные данные
// Вводится натуральное число N, а затем N чисел.

// Выходные данные
// Выведите количество чисел, которые равны нулю.

// Sample Input:
// 5
// 1
// 8
// 100
// 0
// 12
// Sample Output:
// 1

package main

import (
    "fmt"
)

func main() {
    var length int
    fmt.Scan(&length)
    var count int = 0
    var number int
    var i int = 0
    
    for fmt.Scan(&number); i < length; fmt.Scan(&number) {
        if number == 0 {
            count += 1
        }
        i++
    }
    
    fmt.Println(count)
}
