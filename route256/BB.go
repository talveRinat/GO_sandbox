/*
Построение строки

Условия
При приемке посылок на склад кладовщик подписывает и складывает накладные в стопку.
Каждой посылке соответствует ровно одна накладная.
При этом иногда кладовщик переворачивает стопку обратной стороной, но документы всегда кладет на нее сверху.

Требуется составить алгоритм, который поможет кладовщику определить порядок,
в котором сложены документы, если он помнит порядок своих действий.

При этом, пусть каждой накладной присвоена буква латинского алфавита.
Формализованная постановка задачи выглядит так:

Изначально у вас есть пустая строка, далее к ней применяется n модификаций.

Есть два вида модификаций:
Дописать в конец строки символ c
Развернуть строку
Ваша задача вывести итоговую строку после применения всех модификаций.

Формат входных данных
В первой строке входных данных содержится одно целое число n — число модификаций (1 ≤ n ≤ 105).

В следующих n строках содержатся описания модификаций.

«+c», если в конец строки нужно дописать символ c, где символ c — строчный символ латинского алфавита

«!», если строку нужно развернуть


Гарантируется, что итоговая строка состоит хотя бы из одного символа.
Формат выходных данных
Выведите итоговую строку.

Примеры
Входные данные:
5
+a
+b
!
+c
!

Выходные данные:
cab
*/

package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var (
		n         int
		res, text string
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&text)
		ltr := strings.Trim(text, "+")
		if ltr != "!" {
			res += ltr
		} else {
			res = Reverse(res)
		}
	}

	fmt.Println(res)
}
