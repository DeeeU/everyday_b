package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)
	x, max := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			x++
			if x > max {
				max = x
			}
		} else {
			x--
		}
	}
	fmt.Println(max)
}
