package main

import "fmt"

func main() {
	var (
		n, t, sum int
		A         []int
	)
	fmt.Scan(&n, &t)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		A = append(A, num)
	}
	for j := 0; j < n-1; j++ {
		sum += min(t, A[j+1]-A[j])
	}

	fmt.Println(sum + t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
