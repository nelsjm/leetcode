package twosums

import "fmt"

func twoSum(nums []int, target int) []int {
	compliments := make(map[int]int) // [expected compliment] = original index

	for i, val := range nums {
		if idx, ok := compliments[val]; ok {
			fmt.Println("found")
			return []int{idx, i}
		}

		comp := target - val
		compliments[comp] = i
	}

	return []int{}
}
