package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var a, b, c, d int
	a, b, c, d = nextInt(), nextInt(), nextInt(), nextInt()
	fmt.Println(max(max((a*c), (a*d)), max((b*c), (b*d))))
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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
