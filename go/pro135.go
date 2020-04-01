package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
		if m[i] != i+1 {
			count++
		}
	}
	if count <= 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
