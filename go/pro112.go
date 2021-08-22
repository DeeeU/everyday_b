package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inf = int(1e10)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var (
		n, T int
		c, t int
	)

	n, T = nextInt(), nextInt()

	cost := inf

	for i := 0; i < n; i++ {
		c, t = nextInt(), nextInt()
		if cost > c && T >= t {
			cost = c
		}
	}
	if cost != inf {
		fmt.Println(cost)
	} else {
		fmt.Println("TLE")
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
