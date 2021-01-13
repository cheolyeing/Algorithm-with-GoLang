package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var cnt int = 0
	for i := 1; i <= n; i++ {
		var str string = strconv.Itoa(i)

		if len(str) == 1 {
			cnt++
		} else {
			var first_gap int = int(str[0] - str[1])
			var add bool = true

			for j := 2; j < len(str); j++ {
				var next_gap int = int(str[j-1] - str[j])
				if first_gap != next_gap {
					add = false
					break
				}
			}

			if add {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
