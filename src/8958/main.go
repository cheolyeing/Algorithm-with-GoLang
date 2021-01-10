package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var record string
		fmt.Scanln(&record)
		printScore(record)
	}
}

func printScore(record string) {
	var total int = 0
	var point int = 1
	for _, e := range record {
		//fmt.Println(i, string(e))
		if string(e) == "O" {
			total += point
			point++
		} else {
			point = 1
		}
	}
	fmt.Println(total)
}
