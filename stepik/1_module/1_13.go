/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
8

Sample Output:
6
*/

package main

import "fmt"

func main() {
	var (
		a             int          // вводимое число
		f1, f2, fn, n = 0, 1, 0, 1 // 1-е, 2-ое, f_n  число фибоначчи и счетчик
	)

	fmt.Scan(&a)

	// вычисляем число фионначи до водимого число и увеличивыем счетчик
	for fn < a {
		fn = f1 + f2
		f1, f2 = f2, fn
		n++
	}

	if fn == a {
		fmt.Print(n)
	} else {
		fmt.Print(-1)
	}
}