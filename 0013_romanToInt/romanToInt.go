package romanToInt

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

var lookup = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	pos := 0
	val := 0
	for pos < len(s) {
		if pos+1 < len(s) {
			sub := s[pos : pos+2]
			if v, ok := lookup[sub]; ok {
				val += v
				pos += 2
				continue
			}
		}

		sub := string(s[pos : pos+1])

		if v, ok := lookup[sub]; ok {
			val += v
			pos += 1
			continue
		}
	}

	return val
}
