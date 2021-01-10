package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	for i := 2; i >= 0; i-- {
		var n = b[i] - '0'
		fmt.Println(a * int(n))
	}
	intStr, _ := strconv.ParseInt(b, 10, 64)
	fmt.Println(a * int(intStr))
}
