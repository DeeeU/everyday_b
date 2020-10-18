package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N int
	N = nextInt()
	nl := make([]int, N)
	for i := 0; i < N; i++ {
		nl[i] = nextInt()
	}
	fmt.Println(mn(nl))
	fmt.Println(y(nl))
	fmt.Println(tye(nl))
}

func init() {
	const max = 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
	sc.Split(bufio.ScanWords)
	return
}

func mn(s []int) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] < 0 {
			s[i] = s[i] * -1
		}
		count += s[i]
	}
	return count
}

func y(s []int) float64 {
	count := 0
	for i := 0; i < len(s); i++ {
		count += s[i] * s[i]
	}
	return math.Sqrt(float64(count))
}

func tye(s []int) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if count < s[i] {
			count = s[i]
		}
	}
	return count
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
