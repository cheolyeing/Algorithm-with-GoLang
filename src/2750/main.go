package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	sort.Ints(arr)

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
