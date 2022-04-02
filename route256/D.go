/*
Оригинальный пароль

Условия
Василий постоянно забывает не только ПИН-коды от банковских карт, но и пароли от различных онлайн-сервисов. Чтобы победить свою забывчивость, он решил выработать для себя такое правило придумывания паролей: каждый пароль должен состоять из четного количества символов, и при этом первая половина пароля должна быть «анаграммно меньше» второй его половины. Теперь Василию нужен алгоритм, проверяющий, удовлетворяет ли придуманный пароль заданному условию.
Вам дано n пар строк одинаковой длины (каждая строка – половина пароля).

Строка s считается анаграммно меньше строки t, если существуют строка s' и строка t' такие что:
s' получается из s перестановкой букв
t' получается из t перестановкой букв
s' лексикографически меньше, чем t'

Для каждой пары строк si, ti определите правда ли, что si анаграммно меньше чем ti.

Формат входных данных
В первой строке задано целое число n - число пар строк в тесте (1 < n < 100).

В следующих 3n строках содержатся описания пар строк. Каждое описание состоит из трех последовательных строк. В первой строке каждого описания дана длина строк m_i (1 < m_i < 100), затем в следующих двух строках записаны строки  s_i, t_i длины m_i, состоящие из строчных латинских букв.

Формат выходных данных
Для каждой пары строк из входных данных, выведите в i-й строке Yes, если строка s_i анаграммно меньше строки t_i. Иначе выведите No.

Примеры
Входные данные:
3
3
cba
aaz
2
bc
ab
1
a
a
Выходные данные:
Yes
No
No
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readStdInputSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	var i int
	for scanner.Scan() {
		i++
		input = append(input, scanner.Text())
		numberOfPasswords, _ := strconv.Atoi(input[0])
		if numberOfPasswords*3 == i-1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return input
}

func originalPassword(inputData []string) []string {
	var answer []string
	numberOfPasswords, _ := strconv.Atoi(inputData[0])
	for i := 0; i < numberOfPasswords*3; i = i + 3 {
		if compareStrings(inputData[i+2], inputData[i+3]) {
			answer = append(answer, "YES")
		} else {
			answer = append(answer, "NO")
		}
	}
	return answer
}

func compareStrings(firstHalf, secondHalf string) bool {
	for _, firstHalfLetter := range firstHalf {
		for _, secondHalfLetter := range secondHalf {
			if firstHalfLetter < secondHalfLetter {
				return true
			}
		}
	}
	return false
}
func writeStdOutputSlice(output []string) {
	for _, val := range output {
		fmt.Fprint(os.Stdout, val+"\n")
	}
}
func main() {
	input := readStdInputSlice()
	passwordChecks := originalPassword(input)
	writeStdOutputSlice(passwordChecks)
}
