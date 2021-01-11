package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var MAX int = -1000000
	var MIN int = 1000000

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")

	nums := strings.Split(line, " ")

	for _, elem := range nums {
		num, _ := strconv.Atoi(elem)

		if num > MAX {
			MAX = num
		}
		if num < MIN {
			MIN = num
		}
	}

	fmt.Println(MIN, MAX)
}
