/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и
отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
во второй вы должны отправлять значения без повторов.
В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
*/
package main

import "fmt"

func removeDuplicates(input chan string, output chan string) {
	var prev string
	defer close(output)

	for item := range input {
		if prev != item {
			output <- item
		}
		prev = item
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	// 123456
}
