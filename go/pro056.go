package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)
	if a+w < b {
		fmt.Print(b - (a + w))
	} else if b+w < a {
		fmt.Print(a - (b + w))
	} else {
		fmt.Print(0)
	}
}
