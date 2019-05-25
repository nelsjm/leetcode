package intToRoman

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	roman := ""
	for num > 0 {
		if num >= 900 {
			roman, num = largerNumeral(roman, num, 1000, "M", 100, "C")
			continue
		}

		if num >= 400 {
			roman, num = largerNumeral(roman, num, 500, "D", 100, "C")
			continue
		}

		if num >= 90 {
			roman, num = largerNumeral(roman, num, 100, "C", 10, "X")
			continue
		}

		if num >= 40 {
			roman, num = largerNumeral(roman, num, 50, "L", 10, "X")
			continue
		}

		if num >= 9 {
			roman, num = largerNumeral(roman, num, 10, "X", 1, "I")
			continue
		}

		if num >= 4 {
			roman, num = largerNumeral(roman, num, 5, "V", 1, "I")
			continue
		}

		roman += "I"
		num--
	}

	return roman
}

func largerNumeral(curr string, num int, base int, baseNumeral string, modifier int, modifierNumeral string) (string, int) {
	if num >= base {
		curr += baseNumeral
		num -= base
		return curr, num
	}

	if num%base >= base-modifier {
		curr += modifierNumeral + baseNumeral
		return curr, num - (base - modifier)
	}

	return curr, num
}
