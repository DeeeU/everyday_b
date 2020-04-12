package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	sum, max := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		if max < m {
			max = m
		}
		sum += m
	}
	if max < sum-max {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
