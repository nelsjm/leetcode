package longestSubstring

import "testing"

func Test_LongestSubstring1(t *testing.T) {
	tc := "abcabcbb"
	expected := 3

	res := lengthOfLongestSubstring(tc)
	if res != expected {
		t.Errorf("expected length %v but got %v", expected, res)
	}
}

func Test_LongestSubstring2(t *testing.T) {
	tc := "bbbbb"
	expected := 1

	res := lengthOfLongestSubstring(tc)
	if res != expected {
		t.Errorf("expected length %v but got %v", expected, res)
	}
}

func Test_LongestSubstring3(t *testing.T) {
	tc := "pwwkew"
	expected := 3

	res := lengthOfLongestSubstring(tc)
	if res != expected {
		t.Errorf("expected length %v but got %v", expected, res)
	}
}

func Test_LongestSubstringAllUnique(t *testing.T) {
	tc := "abcdefg"
	expected := 7

	res := lengthOfLongestSubstring(tc)
	if res != expected {
		t.Errorf("expected length %v but got %v", expected, res)
	}
}
