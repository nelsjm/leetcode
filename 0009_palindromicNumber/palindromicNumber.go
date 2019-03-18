package palindromicNumber

import "math"

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	v := x
	reversed := 0
	for v != 0 {
		reversed = (reversed * 10) + (v % 10)
		v = int(math.Floor(float64(v / 10)))
	}

	return x == reversed
}
