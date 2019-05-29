package threeSums

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })

	results := make(map[string][]int, 0)

	for i := range nums {
		a := nums[i]

		j := i + 1
		k := len(nums) - 1

		for j < k {
			total := a + nums[j] + nums[k]
			if total == 0 {
				results[fmt.Sprintf("%v|%v|%v|", a, nums[j], nums[k])] = []int{a, nums[j], nums[k]}
				j++
				continue
			}

			if total < 0 {
				j++
			} else {
				k--
			}

		}
	}

	finalResults := make([][]int, 0, len(results))
	for _, v := range results {
		finalResults = append(finalResults, v)
	}

	return finalResults
}
