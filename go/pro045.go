package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	i := byte('a')
	for {
		if i == 'a' {
			if len(a) == 0 {
				fmt.Println("A")
				return
			}
			i = a[0]
			a = a[1:]
		} else if i == 'b' {
			if len(b) == 0 {
				fmt.Println("B")
				return
			}
			i = b[0]
			b = b[1:]
		} else {
			if len(c) == 0 {
				fmt.Println("C")
				return
			}
			i = c[0]
			c = c[1:]
		}
	}
}
