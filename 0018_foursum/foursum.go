package foursum

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	dedupe := make(map[string][]int)

	for l := 0; l < len(nums)-2; l++ {
		for r := len(nums) - 1; r > l+1; r-- {
			il := l + 1
			ir := r - 1
			sum := 0
			for il < ir {
				sum = nums[l] + nums[il] + nums[ir] + nums[r]
				if sum == target {
					dedupe[fmt.Sprintf("%v|%v|%v|%v", nums[l], nums[il], nums[ir], nums[r])] = []int{nums[l], nums[il], nums[ir], nums[r]}
					il++
					continue
				}

				if sum < target {
					il++
				} else {
					ir--
				}
			}
		}
	}

	results := make([][]int, 0, len(dedupe))
	for _, v := range dedupe {
		results = append(results, v)
	}

	return results
}
