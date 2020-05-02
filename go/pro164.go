package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	m := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	fmt.Println(len(removeDuplicate1(m)))
}

func removeDuplicate1(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}
