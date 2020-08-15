package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type pair struct {
	rn string
	p  int
	i  int
}

func main() {
	n := nextInt()
	s := make([]pair, n)

	for i := 0; i < n; i++ {
		s[i].rn = nextStr()
		s[i].p = nextInt()
		s[i].i = i + 1
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i].rn == s[j].rn {
			return s[i].p > s[j].p
		}
		return s[i].rn < s[j].rn
	})

	for _, v := range s {
		fmt.Println(v.i)
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
