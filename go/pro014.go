package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n, X int
	n, X = nextInt(), nextInt()
	m := make([]int, n)

	count := 0

	for i := 0; i < n; i++ {
		m[i] = nextInt()
		if (1<<i)&X != 0 {
			count += m[i]
		}
	}

	fmt.Println(count)
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
