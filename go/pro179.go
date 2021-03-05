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
		n     int
		count int
		ans   int
	)
	count, ans = 0, 0
	n = nextInt()
	for i := 0; i < n; i++ {
		s1, s2 := nextInt(), nextInt()
		if s1 == s2 {
			count++
		} else {
			count = 0
		}
		if count >= 3 {
			ans = 1
		}
	}
	if ans >= 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
