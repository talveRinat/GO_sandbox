/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().
*/
package main

import "fmt"

func task2(c chan string, text string) {
	for i := 0; i < 5; i++ {
		c <- text + " "
	}
}

func main() {
	c := make(chan string)
	go task2(c, "a")

	for i := 0; i < 5; i++ {
		fmt.Print(<-c)
	}
	// a a a a a
}
