package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	sum := (n - 1) / 2
	ans := "b"
	i := 0
	for i < sum {
		if i%3 == 0 {
			ans = "a" + ans + "c"
		} else if i%3 == 1 {
			ans = "c" + ans + "a"
		} else {
			ans = "b" + ans + "b"
		}
		i++
	}
	if s == ans {
		fmt.Println(sum)
	} else {
		fmt.Println(-1)
	}
}
