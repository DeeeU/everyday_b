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
		n, k, m, sum int
	)

	n, k, m = nextInt(), nextInt(), nextInt()

	a := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		a[i] = nextInt()
		sum += a[i]
	}

	if n*m-sum <= k {
		if n*m-sum <= 0 {
			fmt.Print(0)
		} else {
			fmt.Print(n*m - sum)
		}
	} else {
		fmt.Print("-1")
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
