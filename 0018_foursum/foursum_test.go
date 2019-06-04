package foursum

import (
	"fmt"
	"testing"
)

var cases = []struct {
	in     []int
	target int
	out    []int
}{
	{
		in:     []int{-3, -2, -1, 0, 0, 1, 2, 3},
		target: 0,
	},
}

func Test_fourSum(t *testing.T) {
	for _, c := range cases {
		res := fourSum(c.in, c.target)
		fmt.Println(res)
	}
}
