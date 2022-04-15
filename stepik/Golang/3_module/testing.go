package main

import "fmt"

func specialPrint(arg ...int) error {
	if _, err := fmt.Print(arg); err != nil {
		return err
	}
	return nil
}

func main() {
	specialPrint(10, 4, 5, 8, 7)
}
