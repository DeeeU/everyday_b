package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	w := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
	}
	result := 1
	for i := 1; i < n; i++ {
		if w[i][0] != w[i-1][len(w[i-1])-1] {
			result = 0
		}
		for j := 0; j < i; j++ {
			if w[i] == w[j] {
				result = 0
			}
		}
		if result == 0 {
			break
		}
	}
	if result == 1 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
