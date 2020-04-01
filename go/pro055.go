package main

import "fmt"

func main() {
	var s, n int
	fmt.Scan(&s)
	n = 1
	for i := 1; i < s+1; i++ {
		n *= i
		n %= 1000000007
	}
	fmt.Println(n)
}
