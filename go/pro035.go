package main

import (
	"fmt"
)

func main() {
	x, y, m := 0, 0, 0
	var t int
	var s string
	fmt.Scan(&s)
	fmt.Scan(&t)
	for _, i := range s {
		if i == 'L' {
			x--
		} else if i == 'R' {
			x++
		} else if i == 'U' {
			y++
		} else if i == 'D' {
			y--
		} else {
			m++
		}
	}

	if t == 1 {
		fmt.Println(abs(x) + abs(y) + m)
	} else {
		ans := abs(x) + abs(y)
		if ans >= m {
			fmt.Println(ans - m)
		} else {
			fmt.Println((m - ans) % 2)
		}
	}

}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
