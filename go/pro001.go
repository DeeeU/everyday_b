package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var m, vv int
	m = nextInt()
	if m < 100 {
		vv = 0
	} else if 001 <= m && m <= 5000 {
		m *= 10
		vv = m / 1000
	} else if 6000 <= m && m <= 30000 {
		vv = m/1000 + 50
	} else if 35000 <= m && m <= 70000 {
		vv = ((m/1000 - 30) / 5) + 80
	} else {
		vv = 89
	}
	// %02
	fmt.Printf("%02d\n", vv)
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
