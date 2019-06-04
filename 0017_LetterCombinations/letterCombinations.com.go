package letterCombinations

import "fmt"

var lookup = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return addNumber(digits)
}

func addNumber(digits string) []string {
	if digits == "" {
		return []string{}
	}

	if len(digits) == 1 {
		return lookup[digits]
	}

	set := addNumber(digits[1:])
	digit := string(digits[0])
	letters := lookup[digit]

	results := make([]string, 0)
	for _, l := range letters {
		for _, s := range set {
			results = append(results, fmt.Sprintf("%s%s", l, s))
		}
	}

	return results

}
