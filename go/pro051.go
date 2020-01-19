package main

import (
	"fmt"
)

func main() {
	var k, s int
	ans := 0
	fmt.Scan(&k, &s)
	for i := 0; i <= k; i++ {
		for t := 0; t <= k; t++ {
			z := s - i - t
			if 0 <= z && z <= k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
