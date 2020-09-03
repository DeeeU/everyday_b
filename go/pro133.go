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
	sc.Split(bufio.ScanWords)

	n := nextInt()
	dim := nextInt()
	//var xlist [n][dim]int
	xlist := make([][]int, n)

	for i := 0; i < n; i++ {
		xlist[i] = make([]int, dim)
		for j := 0; j < dim; j++ {
			xlist[i][j] = nextInt()
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}

			tmp := 0
			for di := 0; di < dim; di++ {
				tmp += (xlist[i][di] - xlist[j][di]) * (xlist[i][di] - xlist[j][di])
			}
			dist := math.Sqrt(float64(tmp))

			if isInteger(dist) {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

// ---- readfunc
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
