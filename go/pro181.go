package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	sum := 0
	n := nextInt()
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		sum += (b*(b+1)/2 - a*(a-1)/2)
	}
	fmt.Println(sum)
}

func init() {
	const max = 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
	sc.Split(bufio.ScanWords)
	return
}

func sq(a int) int {
	count := 1
	for i := 1; i < a+1; i++ {
		count += i
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

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
