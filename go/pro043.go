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
		s, ans string
	)

	s = nextStr()
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			if len(ans) != 0 {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans += string(s[i])
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

func remove(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
