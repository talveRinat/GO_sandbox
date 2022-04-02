/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные
Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4

Sample Output:
1

*/

package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	arr := make([]int, n)

	// заполняем массив
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// находим минимальный элемент
	min := arr[0]
	for i := 0; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	// находим сколько раз встречается минимальный элемент
	for i := 0; i < n; i++ {
		if min == arr[i] {
			cnt++
		}
	}
	// вывод
	fmt.Println(cnt)
}