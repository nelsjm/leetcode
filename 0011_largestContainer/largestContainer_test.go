package largestContainer

import (
	"fmt"
	"testing"
)

var cases = []struct {
	values   []int
	expected int
}{
	{
		values:   []int{1, 2},
		expected: 1,
	},
	{
		values:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
}

func Test_MaxArea(t *testing.T) {
	for i, tc := range cases {
		r := maxArea(tc.values)

		if r != tc.expected {
			t.Error(fmt.Sprintf("%v) expected value %v but got %v", i, tc.expected, r))
		}
	}
}
