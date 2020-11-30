package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n, k int
	n, k = nextInt(), nextInt()
	ans, count := 1, 0
	if k != 10 {
		for i := 0; i < n; i++ {
			ans *= k
			count++
			if ans > n {
				break
			}
		}
		fmt.Println(count)
	} else {
		fmt.Println(len(strconv.Itoa(n)))
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

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
