package threesumclosest

import "testing"

func Test_Closest1(t *testing.T) {
	in := []int{-1, 2, 1, -4}
	target := 1
	expects := 2

	res := threeSumClosest(in, target)
	if res != expects {
		t.Errorf("expected %v but got %v", expects, res)
	}
}

func Test_Closest82(t *testing.T) {
	in := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	expects := 82

	res := threeSumClosest(in, target)
	if res != expects {
		t.Errorf("expected %v but got %v", expects, res)
	}
}
