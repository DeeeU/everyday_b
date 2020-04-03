package main

import (
	"fmt"
)

func div(t int) int {
	count := 0
	for i := 1; i <= t; i++ {
		if t%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	c := 0
	for i := 1; i <= n; i += 2 {
		if div(i) == 8 {
			c++
		}
	}
	fmt.Println(c)
}
