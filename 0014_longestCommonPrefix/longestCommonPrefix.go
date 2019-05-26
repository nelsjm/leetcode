package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]

	for _, s := range strs[1:] {
		if s == prefix {
			continue
		}

		newPrefix := ""
		for i := 0; i < min(len(prefix), len(s)); i++ {
			if s[i] == prefix[i] {
				newPrefix += string(s[i])
			} else {
				break
			}
		}

		prefix = newPrefix
		if prefix == "" {
			break
		}
	}

	return prefix
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
