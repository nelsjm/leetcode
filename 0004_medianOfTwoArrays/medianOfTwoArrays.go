package medianOfTwoArrays

import (
	"math"
)

// num1 and num2 will be sorted find the median in O(logn) time
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var small, big []int
	if len(nums1) < len(nums2) {
		small, big = nums1, nums2
	} else {
		small, big = nums2, nums1
	}

	lenSmall := len(small)
	lenBig := len(big)

	low := 0
	high := len(small)

	for low <= high {

		partitionSmall := (low + high) / 2
		partitionBig := (lenSmall+lenBig+1)/2 - partitionSmall

		maxSmallLeft, maxBigLeft := int(math.MinInt32), int(math.MinInt32)
		minSmallRight, minBigRight := int(math.MaxInt32), int(math.MaxInt32)

		if partitionSmall > 0 {
			maxSmallLeft = small[partitionSmall-1]
		}

		if partitionSmall != lenSmall {
			minSmallRight = small[partitionSmall]
		}

		if partitionBig > 0 {
			maxBigLeft = big[partitionBig-1]
		}

		if partitionBig != lenBig {
			minBigRight = big[partitionBig]
		}

		if maxSmallLeft <= minBigRight && minSmallRight >= maxBigLeft {
			if (lenSmall+lenBig)%2 == 0 {
				return (float64(max(maxSmallLeft, maxBigLeft)) + float64(min(minSmallRight, minBigRight))) / 2.0
			} else {
				return float64(max(maxSmallLeft, maxBigLeft))
			}
		}

		if maxSmallLeft > minBigRight {
			high = partitionSmall - 1
		} else {
			low = partitionSmall + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
