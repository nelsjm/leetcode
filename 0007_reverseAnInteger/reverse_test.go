package reverseAnInteger

import "testing"

func Test_PosNumber(t *testing.T) {
	in := 123
	out := 321
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_PosNumberTrailing0(t *testing.T) {
	in := 1230
	out := 321
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_NegativeNumber(t *testing.T) {
	in := -123
	out := -321
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_NegativeNumberTrailing0(t *testing.T) {
	in := -1230
	out := -321
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_NegativeNumberEvenLength(t *testing.T) {
	in := -1234
	out := -4321
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_Zero(t *testing.T) {
	in := 0
	out := 0
	res := reverse(in)

	if res != out {
		t.Errorf("expected value %v but got %v", out, res)
	}
}

func Test_OutofBounds(t *testing.T) {
	in := 1534236469
	out := 0
	res := reverse(in)

	if out != res {
		t.Errorf("expected value %v but got %v", out, res)
	}
}
