package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c int
	slice := make([]int, 10)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	var multi int = a * b * c

	var multiStr string = strconv.Itoa(multi)

	for _, digit := range multiStr {
		slice[digit-'0']++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(slice[i])
	}
}
