/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint),
возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Можно использовать - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:
727178
Sample Output:
28
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(num uint) uint {
		var n string
		str := strconv.FormatUint(uint64(num), 10)
		for i := range str {
			if str[i] == '0' {
				continue
			}
			if (str[i]-'0')%2 == 0 {
				n += string(str[i])

			}
		}
		res, _ := strconv.ParseUint(n, 10, 64)
		if res == 0 {
			return 100
		} else {
			return uint(res)
		}
	}

	fmt.Println(fn(357))
}
