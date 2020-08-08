package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)

	fmt.Scan(&n)
	max := 1000000000000000000
	a := make([]int, n)
	sum := 1
	for i := range a {
		fmt.Scan(&a[i])
		if a[i] == 0 {
			fmt.Print(0)
			return
		}
	}
	for _, j := range a {
		if j <= max/sum {
			sum *= j
		} else {
			fmt.Print(-1)
			return
		}
	}
	fmt.Print(sum)
}
