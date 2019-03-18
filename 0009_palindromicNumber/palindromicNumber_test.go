package palindromicNumber

import "testing"

func Test_Example1(t *testing.T) {
	input := 121
	output := true
	//Explanation 121 is a palindrome

	res := isPalindrome(input)
	if res != output {
		t.Errorf("expected a %v value but got %v", output, res)
	}
}

func Test_Example2(t *testing.T) {
	input := -121
	output := false
	//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

	res := isPalindrome(input)
	if res != output {
		t.Errorf("expected a %v value but got %v", output, res)
	}
}

func Test_Example3(t *testing.T) {
	input := 10
	output := false
	//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

	res := isPalindrome(input)
	if res != output {
		t.Errorf("expected a %v value but got %v", output, res)
	}
}
