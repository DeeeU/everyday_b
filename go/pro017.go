package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	x := nextStr()
	if len(x) == 0 {
		fmt.Println("YES")
		return
	}

	for i := 0; i < len(x); i++ {
		now := x[i]
		if now == 'o' || now == 'k' || now == 'u' {
			continue
		} else if i < len(x)-1 && now == 'c' && x[i+1] == 'h' {
			i++
			continue
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
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
