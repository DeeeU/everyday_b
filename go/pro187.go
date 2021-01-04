package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	ans := 0
	n = nextInt()
	x, y := make([]float64, n), make([]float64, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextFloat64(), nextFloat64()
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ch := (y[j] - y[i]) / (x[j] - x[i])
			if ch <= 1 && ch >= -1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func init() {
	const max = 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
	sc.Split(bufio.ScanWords)
	return
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nextStr() string {
	sc.Scan()
	return sc.Text()
}
func nextFloat64() float64 {
	f, err := strconv.ParseFloat(nextStr(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
