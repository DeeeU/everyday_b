package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scan(&s)
	if n%2 != 0 {
		fmt.Print("No")
	} else if s[0:n/2] == s[n/2:n] {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
