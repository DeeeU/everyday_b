package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	s = nextStr()

	flag := true

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			flag = false
			return
		}
	}
	for i := 0; i < (len(s)-1)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			flag = false
		}
	}
	for i := (len(s) + 3) / 2; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			flag = false
			return
		}
	}
	if flag {
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

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
