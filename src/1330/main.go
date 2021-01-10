package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	var print string
	if a == b {
		print = "=="
	} else {
		if a < b {
			print = "<"
		} else {
			print = ">"
		}
	}

	fmt.Println(print)
}
