package intToRoman

import (
	"fmt"
	"testing"
)

var cases = []struct {
	in  int
	out string
}{
	{in: 0, out: ""},
	{in: 3, out: "III"},
	{in: 4, out: "IV"},
	{in: 9, out: "IX"},
	{in: 10, out: "X"},
	{in: 58, out: "LVIII"},
	{in: 1994, out: "MCMXCIV"},
}

func Test_Cases(t *testing.T) {
	for i, c := range cases {
		out := intToRoman(c.in)
		if out != c.out {
			t.Error(fmt.Sprintf("%v) %v expected %s but got %s", i, c.in, c.out, out))
		}
	}
}
