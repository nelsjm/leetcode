package ubst

import "fmt"

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + f(i,n)
	}

	return sum
}

var cache = make(map[string]int)

func f(i, n int) int {
	key := fmt.Sprintf("%v:%v", i,n)
	if val, ok := cache[key]; ok {
		return val
	}

	v := numTrees(i-1) * numTrees(n-i)
	cache[key] = v
	return v
}