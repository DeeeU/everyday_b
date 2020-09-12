package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == n {
				count++
			}
		}
	}
	if count > 0 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
