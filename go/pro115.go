package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	max, sum := 0, 0
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		if max < t {
			max = t
		}
		sum += t
	}
	fmt.Print(sum - (max / 2))
}
