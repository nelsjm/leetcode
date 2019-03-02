package twosums

import (
	"testing"
)

func Test_TwoSums(t *testing.T) {
	tcNums := []int{2,7,11,15}
	tcTarget := 9
	expects := []int{0,1}

	result := twoSum(tcNums, tcTarget)

	if len(result) != 2 {
		t.Fatalf("expected 2 results but got %v", len(result))
	}

	for i := range expects {
		if result[i] != expects[i] {
			t.Errorf("expected val %v but got %v", expects[i], result[i])
		}
	}
}
