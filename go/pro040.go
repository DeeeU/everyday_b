package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := n
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			r := n - i*j
			if r < 0 {
				break
			}
			v := i - j
			if v < 0 {
				v = -v
			}

			if r+v < ans {
				ans = r + v
			}
		}
	}
	fmt.Println(ans)
}
