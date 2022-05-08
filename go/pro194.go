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
		a, b, n, ma, mb, ans int
	)
	n = nextInt()
	for i := 0; i < n; i++ {
		a, b = nextInt(), nextInt()
		if i == 0 {
			ans = a + b
			ma = a
			mb = b
			continue
		}
		if max(ma, b) < ans {
			ans = max(ma, b)
		}
		if max(mb, a) < ans {
			ans = max(mb, a)
		}
		if a+b < ans {
			ans = a + b
		}
		if ma > a {
			ma = a
		}
		if mb > b {
			mb = b
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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
