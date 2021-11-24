package coinchange

func coinChange(coins []int, amount int) int {
	cache := make([]int, amount+1)
	for i := 0; i < amount + 1; i++ {
		cache[i] = amount + 1
	}

	cache[0] = 0

	for a := 1; a<= amount; a++ {
		for _, c := range coins {
			if a - c >= 0 {
				cache[a] = min(cache[a], 1 + cache[a-c])
			}
		}
	}

	if cache[amount] == amount + 1 {
		return -1
	}

	return cache[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

