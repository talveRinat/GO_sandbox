/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64
(наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:
12 мин. 13 сек.
Sample Output:
Fri May 15 19:28:18 UTC 2020
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
	const now = 1589570165
	startDate := time.Unix(now, 0)

	// считываем данные
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	// заменяем мин. на m, сек. на s и убираем пробелы и \n
	str = strings.Replace(str, "мин.", "m", 1)
	str = strings.Replace(str, "сек.", "s", 1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.TrimRight(str, "\n")

	// тип string заменяем на тип duration
	dur, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}

	// выводим (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате
	fmt.Println(startDate.Add(dur).UTC().Format(time.UnixDate))
}

// короткое решение
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	var m, s time.Duration
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
	fmt.Println(t.Format(time.UnixDate))
}
*/
