package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := n / 3600
	n %= 3600
	m := n / 60
	s := n % 60
	fmt.Printf("%02d:%02d:%02d\n", h, m, s)
}
