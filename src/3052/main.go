package main

import "fmt"

func main() {
	var count int = 0
	slice := make([]bool, 42)

	for i := 0; i < 10; i++ {
		var num int
		fmt.Scanln(&num)

		if !slice[num%42] {
			slice[num%42] = true
			count++
		}
	}

	fmt.Println(count)
}
