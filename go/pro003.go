package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s, t string

	s, t = nextStr(), nextStr()

	word := "atcoder"
	flag := true

	for i := 0; i < len(s); i++ {
		if s[i] == '@' && strings.Contains(word, string(t[i])) {
			continue
		}
		if t[i] == '@' && strings.Contains(word, string(s[i])) {
			continue
		}
		if s[i] != t[i] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("You can win")
	} else {
		fmt.Println("You will lose")
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

func nextFloat64() float64 {
	f, err := strconv.ParseFloat(nextStr(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
