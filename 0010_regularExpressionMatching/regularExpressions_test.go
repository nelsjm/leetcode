package regularExpressionMatching

import "testing"

func Test_Example1(t *testing.T) {
	s := "aa"
	p := "a"
	output := false
	//Explanation: "a" does not match the entire string "aa".

	res := isMatch(s, p)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_Example2(t *testing.T) {
	s := "aa"
	p := "a*"
	output := true
	//Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

	res := isMatch(s, p)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_Example3(t *testing.T) {
	s := "ab"
	p := ".*"
	output := true
	//Explanation: ".*" means "zero or more (*) of any character (.)".

	res := isMatch(s, p)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_Example4(t *testing.T) {
	s := "aab"
	p := "c*a*b"
	output := true
	//Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".

	res := isMatch(s, p)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}

func Test_Example5(t *testing.T) {
	s := "mississippi"
	p := "mis*is*p*."
	output := false
	//Explanation not given

	res := isMatch(s, p)
	if res != output {
		t.Errorf("expected value %v but got %v", output, res)
	}
}
