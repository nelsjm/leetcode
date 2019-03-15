package reverseAnInteger

import "math"

func reverse(x int) int {

	if x == 0 || x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	reversed := 0

	for x != 0 {
		reversed = (reversed * 10) + (x % 10)
		x = int(math.Floor(float64(x / 10)))
	}

	if math.MinInt32 < reversed && reversed < math.MaxInt32 {
		return reversed
	}

	return 0
}
