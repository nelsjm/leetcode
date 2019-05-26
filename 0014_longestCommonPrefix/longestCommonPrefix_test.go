package longestCommonPrefix

import (
	"fmt"
	"testing"
)

var cases = []struct {
	in  []string
	out string
}{
	{in: []string{"aaa", "aa", "a"}, out: "a"},
	{in: []string{"flower", "flow", "flight"}, out: "fl"},
	{in: []string{"dog", "racecar", "car"}, out: ""},
}

func Test_Cases(t *testing.T) {
	for i, c := range cases {
		out := longestCommonPrefix(c.in)
		if out != c.out {
			t.Error(fmt.Sprintf("%v) expected '%s' but got '%s'", i, c.out, out))
		}
	}
}
