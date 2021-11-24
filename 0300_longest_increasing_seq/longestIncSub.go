package longestincsub

func lengthOfLIS(nums []int) int {
	cache := make([]int, len(nums))
	cache[len(nums)-1]=1

	for i := len(nums)-2;i >=0; i = i - 1 {
		vals := []int{1}
		for j := i+1; j < len(nums); j = j + 1 {
			if nums[i] < nums[j] {
				vals = append(vals, 1 + cache[j])
			}
		}
		cache[i]=max(vals)
	}

	return max(cache)
}

func max(a []int) int {
	r := 0
	for _, v := range a {
		if r < v {
			r = v
		}
	}

	return r

}