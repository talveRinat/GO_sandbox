package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/rinatmahmutov/GolandProjects/GO_sandbox/stepik/3_module/task.data")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)
	var count int
	for {
		str, err := r.ReadString(';')
		if err != nil {
			break
		}
		count++
		if str == "0;" {
			fmt.Print(count)
			break
		}
	}
}
