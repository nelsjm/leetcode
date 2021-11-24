package stockwithcooldown


func maxProfit(prices []int) int {
	noStock := make([]int, len(prices))
	ownStock := make([]int, len(prices))
	soldStock := make([]int, len(prices))

	noStock[0] = 0
	ownStock[0] = 0 - prices[0]
	soldStock[0] = 0

	for i := 1; i < len(prices); i = i + 1 {
		// handle state transitions
		noStock[i] = max(noStock[i-1], soldStock[i-1])
		ownStock[i] = max(ownStock[i-1], noStock[i-1]-prices[i])
		soldStock[i] = ownStock[i-1]+prices[i]
	}

	return max(noStock[len(prices)-1], soldStock[len(prices)-1])
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}