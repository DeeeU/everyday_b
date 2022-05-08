package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var (
		a, b, c, d int
	)
	a, b, c, d = nextInt(), nextInt(), nextInt(), nextInt()
	ans1, ans2 := a, 0
	if b < c*d {
		for i := 1; i <= a; i++ {
			ans1 += b
			ans2 += c
			if ans1 <= ans2*d {
				fmt.Println(i)
				break
			}
		}
	} else {
		fmt.Println("-1")
	}
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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
