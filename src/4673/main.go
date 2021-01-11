package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sb strings.Builder
	slice := make([]bool, 10001)

	for i := 1; i <= 10000; i++ {
		if !slice[i] {
			sb.WriteString(strconv.Itoa(i) + "\n")
		}

		var created int = create(i)
		if created < 10001 {
			slice[created] = true
		}
	}
	fmt.Print(sb.String())
}

func create(number int) int {
	var result int = number
	var strNum string = strconv.Itoa(number)

	for _, e := range strNum {
		result += int(e - '0')
	}
	return result
}
