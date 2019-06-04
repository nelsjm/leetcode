package letterCombinations

import (
	"sort"
	"testing"
)

func areSame(l []string, r []string) bool {
	if len(l) != len(r) {
		return false
	}

	sort.SliceStable(l, func(i, j int) bool { return l[i] < l[j] })
	sort.SliceStable(r, func(i, j int) bool { return r[i] < r[j] })

	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}

	return true
}

func Test_LetterCombinations(t *testing.T) {
	in := "23"
	out := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	res := letterCombinations(in)
	if !areSame(out, res) {
		t.Errorf("the expected output did not match the actual output \n %v \n %v", out, res)
	}
}
