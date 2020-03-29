package main

import "fmt"

func main() {
	var m, h, n int
	fmt.Scan(&m, &h, &n)
	var x, y, a, set_x, set_y int
	set_x, set_y = 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y, &a)
		switch a {
		case 1:
			if set_x < x {
				set_x = x
			}
		case 2:
			if x < m {
				m = x
			}
		case 3:
			if set_y < y {
				set_y = y
			}
		case 4:
			if y < h {
				h = y
			}
		}
	}

	if h <= set_y || m <= set_x {
		fmt.Print(0)
	} else {
		fmt.Print((m - set_x) * (h - set_y))
	}
}
