package atoi

import (
	"math"
)

const (
	plus   = 43
	hyphen = 45
	space  = 32
	zero   = 48
	nine   = 57
)

func myAtoi(str string) int {
	numberFound := false
	sign := 0
	resultingNumber := 0

	for _, c := range str {
		if c == space && !numberFound {
			if !numberFound && sign == 0 {
				continue
			}

			return 0
		}

		if c == hyphen && !numberFound {
			if sign == 0 {
				sign = -1
				continue
			}
			return 0
		}

		if c == plus && !numberFound {
			if sign == 0 {
				sign = 1
				continue
			}
			return 0
		}

		if c < zero || c > nine {
			if !numberFound {
				return 0
			}
			break
		}

		numberFound = true
		v := (resultingNumber * 10) + (int(c) - zero)
		if v < resultingNumber {
			if sign >= 0 {
				return math.MaxInt32
			}

			return math.MinInt32
		}
		resultingNumber = v
	}

	if !numberFound {
		return 0
	}

	if sign == 0 {
		sign = 1
	}

	resultingNumber = resultingNumber * sign
	if resultingNumber < math.MinInt32 {
		return math.MinInt32
	}

	if resultingNumber > math.MaxInt32 {
		return math.MaxInt32
	}

	return resultingNumber
}
