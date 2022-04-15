/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/

package main

import "fmt"

func main() {
	var n, v int
	r := make([]int, n)
	for fmt.Scan(&n); n > 0; n /= 10 {
		r = append(r, n%10)
	}
	fmt.Scan(&v)
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] != v {
			fmt.Print(r[i])
		}
	}
}

// без массива
/*
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    result := 0
    for i := 1; a > 0; {
        if a % 10 != b {
            result += a % 10 * i
            i *= 10
        }
        a /= 10
    }
    fmt.Println(result)
}
*/
