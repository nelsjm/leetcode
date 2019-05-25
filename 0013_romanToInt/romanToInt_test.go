package romanToInt

import (
	"fmt"
	"testing"
)

var cases = []struct {
	in  string
	out int
}{
	{out: 0, in: ""},
	{out: 3, in: "III"},
	{out: 4, in: "IV"},
	{out: 9, in: "IX"},
	{out: 10, in: "X"},
	{out: 58, in: "LVIII"},
	{out: 1994, in: "MCMXCIV"},
	{out: 1884, in: "MDCCCLXXXIV"},
}

func Test_Cases(t *testing.T) {
	for i, c := range cases {
		out := romanToInt(c.in)
		if out != c.out {
			t.Error(fmt.Sprintf("%v) %s xpected %v but got %v", i, c.in, c.out, out))
		}
	}
}
