package main

import (
	"fmt"
)

func main() {
	var n, X int
	fmt.Scan(&n, &X)
	m := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
		if X>>uint(i)&1 == 1 {
			count += m[i]
		}
	}
	fmt.Println(count)
}
