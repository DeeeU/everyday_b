package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n, m int
	n, m = nextInt(), nextInt()
	s := make([]int, m)
	p := make([]int, m)

	for i := 0; i < m; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			s[j] = nextInt()
		}
	}

	for i := 0; i < m; i++ {
		p[i] = nextInt()
	}

	ans := 0
	for bit := 0; bit < (1 << n); bit++ {
		count := 0
		for j := 0; j < m; j++ {
			if (bit&s[j])%2 == p[j] {
				count++
			}
		}
		if count == m {
			ans++
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

func popcnt(x int) int {
	x = x - x>>1&0x5555555555555555
	x = x&0x3333333333333333 + x>>2&0x3333333333333333
	x = (x + x>>4) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return x & 0x0000007f
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
