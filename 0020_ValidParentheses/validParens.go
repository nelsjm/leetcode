package validParentheses

func isValid(s string) bool {
	pos := -1
	expectedClosers := make([]rune, len(s))

	for _, c := range s {
		if c == '(' {
			pos++
			expectedClosers[pos] = ')'
			continue
		}

		if c == '{' {
			pos++
			expectedClosers[pos] = '}'
			continue
		}

		if c == '[' {
			pos++
			expectedClosers[pos] = ']'
			continue
		}

		if c == ')' || c == '}' || c == ']' {
			if pos < 0 || expectedClosers[pos] != c {
				return false
			}

			pos--
			continue
		}

		return false
	}

	return pos < 0
}
