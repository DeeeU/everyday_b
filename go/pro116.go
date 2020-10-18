package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s int
	s = nextInt()
	slist := make([]int, 1000000)
	for i := 1; i < 1000000; i++ {
		if i == 1 {
			slist[i] = s
		} else if slist[i-1]%2 == 0 {
			slist[i] = slist[i-1] / 2
		} else {
			slist[i] = 3*slist[i-1] + 1
		}
		if slist[i] == 4 || slist[i] == 2 || slist[i] == 1 {
			fmt.Print(i + 3)
			break
		}

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
