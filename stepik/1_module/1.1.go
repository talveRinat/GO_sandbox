package main

/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:
745

Sample Output:
16
*/

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)
}
