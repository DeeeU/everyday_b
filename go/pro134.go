package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	if n%(1+d*2) == 0 {
		fmt.Print(n / (1 + d*2))
	} else {
		fmt.Print(n/(1+d*2) + 1)
	}
}
