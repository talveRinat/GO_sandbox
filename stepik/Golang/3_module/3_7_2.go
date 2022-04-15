/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно,
достаточно вывести дату на стандартный вывод в том же формате.

Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
а затем вывести на стандартный вывод в том же формате.

Sample Input:

2020-05-15 08:00:00
Sample Output:

2020-05-15 08:00:00
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// считываем дату
	rd := bufio.NewReader(os.Stdin)
	timeStr, err1 := rd.ReadString('\n')
	if err1 != nil {
		panic(err1)
	}

	// Удаление последнего символа при помощи пакета "strings".
	timeStr = strings.TrimSuffix(timeStr, "\n")

	// установим шаблон вывода
	format := "2006-01-02 15:04:05"

	timeStruct, err2 := time.Parse(format, timeStr)
	if err2 != nil {
		panic(err2)
	}

	if timeStruct.Hour() < 13 {
		fmt.Println(timeStruct.Format(format))
	} else {
		future := timeStruct.Add(time.Hour * 24)
		fmt.Println(future.Format(format))
	}
}
