package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	slice := make([]string, n)

	for i := 1; i <= n; i++ {
		slice[i-1] = strconv.Itoa(i)
	}

	fmt.Print(strings.Join(slice, "\n"))
}
