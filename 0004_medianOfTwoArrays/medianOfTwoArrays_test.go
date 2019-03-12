package medianOfTwoArrays

import "testing"

func Test_SingleDigits(t *testing.T) {
	v1 := []int{1}
	v2 := []int{3}
	expects := 2.0

	res := findMedianSortedArrays(v1, v2)

	if res != expects {
		t.Errorf("expected value %v but got %v", expects, res)
	}
}

func Test_MultipleDigitsOneGTOther(t *testing.T) {
	v1 := []int{1, 2}
	v2 := []int{3, 4}
	expects := 2.5

	res := findMedianSortedArrays(v1, v2)

	if res != expects {
		t.Errorf("expected value %v but got %v", expects, res)
	}
}

func Test_OddDigits(t *testing.T) {
	v1 := []int{1, 2}
	v2 := []int{3}
	expects := 2.0

	res := findMedianSortedArrays(v1, v2)
	if res != expects {
		t.Errorf("expected value %v but got %v", expects, res)
	}
}

func Test_LargerArrayFirst(t *testing.T) {
	v1 := []int{2, 3, 8}
	v2 := []int{1, 2, 3}
	expects := 2.5

	res := findMedianSortedArrays(v1, v2)
	if res != expects {
		t.Errorf("expected value %v but got %v", expects, res)
	}
}

func Test_MixedNumbers(t *testing.T) {
	v1 := []int{2, 2, 4, 6, 8, 15}
	v2 := []int{3, 6, 7, 9, 10, 12}
	expects := 6.5

	res := findMedianSortedArrays(v1, v2)
	if res != expects {
		t.Errorf("expected value %v but got %v", expects, res)
	}
}
