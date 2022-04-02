/*

Забытый ПИН-код

Условия
Василий забыл ПИН-код от своей банковской карты. Точно помнит только первую цифру.
Остальные три цифры он тоже вспомнил, но забыл в какой последовательности они должны идти.
Василий уверен, что эти три цифры точно разные и среди них нет нуля.
Помогите ему подобрать ПИН-код – выведите на экран все возможные комбинации трёхзначных чисел,
собранные из цифр исходного натурального числа n, в порядке возрастания.

Формат входных данных
Целое число n.
123 ≤ n ≤ 987

Формат выходных данных
Все возможные комбинации в порядке возрастания. Каждая комбинация - на отдельной строке.

Примеры
Входные данные:
123
Выходные данные:
123
132
213
231
312
321

Входные данные:
582
Выходные данные:
258
285
528
582
825
852
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(a string) []string {
	var n func(a []rune, p []string) []string
	n = func(a []rune, p []string) []string {
		if len(a) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), a[0])...)
			}
			return n(a[1:], result)
		}
	}

	output := []rune(a)
	return n(output[1:], []string{string(output[0])})
}

func main() {
	var num int
	fmt.Scan(&num)

	t := strconv.Itoa(num)
	d := permutations(t)
	sort.Strings(d)

	for i := 0; i < len(d); i++ {
		var j int
		if _, err := fmt.Sscan(d[i], &j); err == nil {
			fmt.Printf("%d\n", j)
		}
	}
}
