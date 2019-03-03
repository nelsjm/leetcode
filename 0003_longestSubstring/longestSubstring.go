package longestSubstring

func lengthOfLongestSubstring(s string) int {
	currentLongest := 0
	currentCount := 0

	chars := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if idx, ok := chars[r]; ok {
			if currentCount > currentLongest {
				currentLongest = currentCount
			}

			i = idx
			chars = make(map[rune]int)
			currentCount = 0
			continue
		}

		chars[r] = i
		currentCount++
	}

	if currentCount > currentLongest {
		currentLongest = currentCount
	}

	return currentLongest
}
