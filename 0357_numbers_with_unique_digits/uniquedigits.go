package uniquedigits


func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	if n == 2 {
		return 91
	}

	curr := 91
	prod := 81
	for i := 3; i <=n; i++ {
		prod = prod * (11-i)
		curr = curr + prod
	}

	return curr
}