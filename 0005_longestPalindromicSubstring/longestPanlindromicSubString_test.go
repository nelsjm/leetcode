package longestPalindromicSubstring

import "testing"

func TestIsPalindrom(t *testing.T) {
	memo := make(map[string]bool)

	if !isPalindrome("a", memo) {
		t.Errorf("should be palindrome")
	}

	if !isPalindrome("aa", memo) {
		t.Errorf("should be palindrome")
	}

	if !isPalindrome("aaa", memo) {
		t.Errorf("should be palindrome")
	}

	if !isPalindrome("aaaaa", memo) {
		t.Errorf("should be palindrome")
	}

	if isPalindrome("abc", memo) {
		t.Errorf("should not be palindrome")
	}

	if isPalindrome("abca", memo) {
		t.Errorf("should not be palindrome")
	}

	if !isPalindrome("abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba", memo) {
		t.Errorf("should be a palindrome")
	}
}

var testCases = map[string][]string{
	"babad":                                {"bab", "aba"},
	"cbba":                                 {"bb"},
	"jhsdlfkhjasfuefkslankjsdnfwifecnabba": {"abba"},
	"abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba": {"abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"},
	"abbajhsdlfkhjasfuefkslankjsdnfwifecn":                 {"abba"},
	"jhsdlfkhjasabbafuefkslankjsdnfwifecn":                 {"abba"},
	"ac":                                                   {"a", "c"},
}

func Test_LongestPalindromicSubstring(t *testing.T) {
	for k, vs := range testCases {
		res := longestPalindrome(k)

		found := false
		for _, v := range vs {
			if v == res {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("could not find a valid palindromic substring for '%s' returned '%s'", k, res)
		} else {
			t.Logf("longest palindromic substring for '%s' is '%s'", k, res)
		}
	}
}
