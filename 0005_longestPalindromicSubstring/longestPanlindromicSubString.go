package longestPalindromicSubstring

type startPosition struct {
	left  int
	right int
}

func longestPalindrome(s string) string {
	// finds palindromic centers ie "bb" and "bab"
	positions := make([]startPosition, 0)
	for i := 0; i < len(s); i++ {
		charAtI := s[i]
		ip1 := i + 1
		ip2 := ip1 + 1

		if ip1 >= len(s) {
			continue
		}

		if charAtI == s[ip1] {
			positions = append(positions, startPosition{left: i, right: ip1})
		}

		if ip2 >= len(s) {
			continue
		}

		if charAtI == s[ip2] {
			positions = append(positions, startPosition{left: i, right: ip2})
		}
	}

	if len(positions) == 0 {
		return s[0:1]
	}

	memo := make(map[string]bool)
	currentLongest := ""

	for _, p := range positions {
		res := expandPalindrome(s, p, memo)
		if len(res) > len(currentLongest) {
			currentLongest = res
		}
	}

	return currentLongest
}

func expandPalindrome(s string, position startPosition, memo map[string]bool) string {
	left, right := position.left, position.right+1
	lastSubstring := s[left:right]
	currentSubstring := lastSubstring

	for isPalindrome(currentSubstring, memo) {
		left--
		if left < 0 {
			return currentSubstring
		}

		right++
		if right > len(s) {
			return currentSubstring
		}

		lastSubstring = currentSubstring
		currentSubstring = s[left:right]
	}

	return lastSubstring
}

func isPalindrome(s string, memo map[string]bool) bool {
	is, ok := memo[s]
	if ok {
		return is
	}

	pref, suff := s[0], s[len(s)-1]
	if pref != suff {
		memo[s] = false
		return false
	}

	if len(s) <= 3 {
		memo[s] = true
		return true
	}

	middleIsPalindrome := isPalindrome(s[1:len(s)-1], memo)
	memo[s] = middleIsPalindrome
	return middleIsPalindrome
}
