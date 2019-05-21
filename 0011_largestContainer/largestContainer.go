package largestContainer

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	currentLargest := 0
	left := 0
	right := len(height) - 1

	for left < right {
		lh := height[left]
		rh := height[right]

		minHeight := min(lh, rh)

		size := minHeight * (right - left)
		if size > currentLargest {
			currentLargest = size
		}

		if lh == minHeight {
			left++
		} else {
			right--
		}
	}

	return currentLargest
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
