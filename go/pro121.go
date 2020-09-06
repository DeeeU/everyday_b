package main

import "fmt"

func main() {
	var n, m, a, b, c int
	fmt.Scan(&n, &m, &c)
	sum, count := 0, 0
	gm := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b)
		gm[i] = b
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&a)
			sum = sum + a*gm[j]
		}
		if sum+c > 0 {
			count++
			sum = 0
		}
		sum = 0
	}
	fmt.Print(count)
}
