package maxprodsubarray

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := smax(nums)
	currMin, currMax := 1, 1

	for _, v := range nums {
		if v == 0 {
			currMin, currMax = 1, 1
			continue
		}

		tmp := currMax * v
		currMax = max(v, currMax * v, currMin * v)
		currMin = min(v, tmp, currMin * v)

		res = max(res,res,currMax)
	}

	return res
}

func min(a,b,c int) int {
	if a < b && a < c {
		return a
	}

	if b < c {
		return b
	}

	return c
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}

	if b > c {
		return b
	}

	return c
}

func smax(s []int) int {
	m := math.MinInt32
	for _, v := range s {
		if v > m {
			m = v
		}
	}

	return m
}