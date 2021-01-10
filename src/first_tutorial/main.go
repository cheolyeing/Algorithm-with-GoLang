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

	var MIN int = 1000000
	var MAX int = 1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	slice := strings.Split(input, " ")

	for _, str := range slice {
		num, _ := strconv.Atoi(str)
		fmt.Println(num)

		if MAX < num {
			MAX = num
		}
		if MIN > num {
			MIN = num
		}
	}
	fmt.Println(MIN, MAX)
}
