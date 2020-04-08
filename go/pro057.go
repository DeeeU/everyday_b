package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	a, b := make([]int, N), make([]int, N)
	c, d := make([]int, M), make([]int, M)

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&c[i], &d[i])
	}

	res := make([]int, N)
	for i := 0; i < N; i++ {
		minValue := 2147483647
		minIndex := 0
		for j := 0; j < M; j++ {
			v := abs(a[i]-c[j]) + abs(b[i]-d[j])
			if minValue > v {
				minValue = v
				minIndex = j
			}
		}
		res[i] = minIndex + 1
	}

	for i := 0; i < N; i++ {
		fmt.Println(res[i])
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
