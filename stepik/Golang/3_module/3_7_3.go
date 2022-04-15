/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time,
а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// считываем и проверяем строку
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	//Разделяем строку на подстроки по разделителю " , "
	newStr := strings.Split(str, ",")

	data1, data2 := newStr[0], newStr[1]

	data1 = strings.TrimRight(data1, "\n")
	data2 = strings.TrimRight(data2, "\n")

	// установим шаблон вывода
	format := "02.01.2006 15:04:05"

	// Парсим первую дату
	dataStruct1, err := time.Parse(format, data1)
	if err != nil {
		panic(err)
	}

	// парсим вторую дату
	dataStruct2, err := time.Parse(format, data2)
	if err != nil {
		panic(err)
	}

	// вычислим период между текущим моментом и заданными временами
	t1 := time.Since(dataStruct1).Round(time.Second)
	t2 := time.Since(dataStruct2).Round(time.Second)

	// вывести продолжительность периода между меньшей и большей датами
	if t1 > t2 {
		fmt.Println(dataStruct2.Sub(dataStruct1))
	} else {
		fmt.Println(dataStruct1.Sub(dataStruct2))
	}

}
