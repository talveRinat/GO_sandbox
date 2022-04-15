/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work()
и дождаться результатов ее выполнения.
*/

package main

import "fmt"

func main() {

	work := func() {
		fmt.Println("Hello World!")
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		work()
	}()
	<-done
}
