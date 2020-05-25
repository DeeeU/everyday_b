package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print(n - 1)
	} else {
		fmt.Print(n + 1)
	}
}
