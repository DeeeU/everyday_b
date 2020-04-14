// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	ans := 0.0
// 	for i := 0; i < n; i++ {
// 		var (
// 			x float64
// 			u string
// 		)
// 		fmt.Scan(&x, &u)
// 		if u == "BTC" {
// 			x *= 380000.0
// 		}
// 		ans += x
// 	}
// 	fmt.Println(ans)
// }

package main

import (
	"fmt"
)

func main() {
	var (
		n        int
		money    float64
		currency string
	)

	fmt.Scan(&n)
	total := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&money, &currency)
		if currency == "BTC" {
			money *= 380000.0
		}
		total += money
	}

	fmt.Println(total)
}
