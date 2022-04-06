/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:
843

Sample Output:
348
*/
package main

import "fmt"

func main() {
	var n int

	r := make([]int, n)

	for fmt.Scan(&n); n > 0; n /= 10 {
		r = append(r, n%10)
	}
	fmt.Printf("%d%d%d", r[0], r[1], r[2])
}

/*
package main
import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    for ;n > 0; n /= 10{
        fmt.Print(n % 10)
    }
}
*/
