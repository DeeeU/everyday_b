package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// didn't understand the multiple array
func main() {
	h, w := nextInt(), nextInt()
	g := make([][]string, h)
	for i := 0; i < h; i++ {
		a := nextStr()
		g[i] = make([]string, w)
		for j := 0; j < w; j++ {
			g[i][j] = string(a[j])
		}
	}
	var r, c []int
	// 行を参照
	for i := 0; i < h; i++ {
		flag := true
		for j := 0; j < w; j++ {
			if g[i][j] != "." {
				flag = false
				break
			}
		}
		if !flag {
			r = append(r, i)
		}
	}
	// 列を参照
	for i := 0; i < w; i++ {
		flag := true
		for j := 0; j < h; j++ {
			if g[j][i] != "." {
				flag = false
				break
			}
		}
		if !flag {
			c = append(c, i)
		}
	}
	fmt.Println(c, r)
	for _, i := range r {
		for _, j := range c {
			fmt.Print(g[i][j])
		}
		fmt.Println("")
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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
