package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	s = nextStr()
	flag := true
	for i := range s {
		c := s[i : i+1]
		if i%2 == 0 {
			if c == strings.ToUpper(c) {
				flag = false
			}
		} else {
			if c == strings.ToLower(c) {
				flag = false
			}
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

func nextFloat64() float64 {
	f, err := strconv.ParseFloat(nextStr(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func IsFirstUpper(v string) bool {
	if v == "" {
		return false
	}
	r := rune(v[0])
	return unicode.IsUpper(r)
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
