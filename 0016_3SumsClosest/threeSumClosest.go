package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })

	closest := math.MaxInt32
	diff := math.MaxInt32

	for i := range nums {
		a := nums[i]

		j := i + 1
		k := len(nums) - 1

		for j < k {
			total := a + nums[j] + nums[k]
			currDiff := int(math.Abs(float64(target - total)))

			if currDiff == 0 {
				return target
			}

			if currDiff < diff {
				diff = currDiff
				closest = total
			}

			if total < target {
				j++
			} else {
				k--
			}

		}
	}

	return closest
}
