package atoi

import (
	"math"
	"testing"
)

func Test_TestCase1(t *testing.T) {
	input := "42"
	output := 42
	//Explanation pretty self explanatory

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCase2(t *testing.T) {
	input := "   -42"
	output := -42
	//Explanation: The first non-whitespace character is '-', which is the minus sign. Then take as many numerical digits as possible, which gets 42.

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCase3(t *testing.T) {
	input := "4193 with words"
	output := 4193
	//Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCase4(t *testing.T) {
	input := "words and 987"
	output := 0
	//Explanation: The first non-whitespace character is 'w', which is not a numerical digit or a +/- sign. Therefore no valid conversion could be performed.

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCase5(t *testing.T) {
	input := "-91283472332"
	output := -2147483648
	//Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer. Thefore INT_MIN (−231) is returned.

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCasePI(t *testing.T) {
	input := "3.1415"
	output := 3
	//Explanation: I guess I should stop after finding a nonvalid char ?

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCaseMultipleSigns(t *testing.T) {
	input := "+-2"
	output := 0
	//Explanation: there should only be one sign

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCaseHugeNumber(t *testing.T) {
	input := "9223372036854775808"
	output := math.MaxInt32
	//Explanation: Larger than a valid int 32

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_TestCaseAnotherBigNumber(t *testing.T) {
	input := "18446744073709551617"
	output := math.MaxInt32

	//Explanation: Larger than a valid int 32

	res := myAtoi(input)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}
