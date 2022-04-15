/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            // ...
        }
    ]
}
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

{
    "Average": 14.1
}
Как вы понимаете, для декодирования используется функция Unmarshal,
а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).


*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	Student struct {
		Rating []int
	}
	Group struct {
		Students []Student
	}
)

func main() {
	students := new(Group)

	bytes, _ := ioutil.ReadAll(os.Stdin)
	_ = json.Unmarshal(bytes, students)

	var cnt int
	for _, student := range students.Students {
		cnt += len(student.Rating)
	}

	average := float64(cnt) / float64(len(students.Students))

	bytes, _ = json.MarshalIndent(map[string]float64{"Average": average}, "", "    ")

	fmt.Printf("%s", bytes)
}
