/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7,
или выведите "NO" - если таковых нет.


Sample Input:
100
500

Sample Output:
497
*/

package main

import "fmt"

func main() {
	var a, b, big int
	fmt.Scan(&a, &b)

	for ; a <= b; a++ {
		if a%7 == 0 {
			big = a
		}
	}
	if big == 0 && a != 0 && b != 0 {
		fmt.Println("NO")
	} else if big == 0 && a == 0 && b == 0 {
		fmt.Print(0)
	} else {
		fmt.Println(big)
	}
}
