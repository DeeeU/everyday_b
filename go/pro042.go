package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var (
		n, l int
		s    []string
		tx   string
	)
	fmt.Scan(&n, &l)
	for i := 0; i < n; i++ {
		fmt.Scan(&tx)
		s = append(s, tx)
	}
	sort.Strings(s)
	fmt.Print(strings.Join(s, ""))
}
