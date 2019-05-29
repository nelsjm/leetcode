package threeSums

import (
	"fmt"
	"testing"
)

func Test_Cases(t *testing.T) {
	in := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(in)
	fmt.Println(len(res))
	for _, r := range res {
		for _, i := range r {
			fmt.Print(fmt.Sprintf("%v\t", i))
		}
		fmt.Print("\n")
	}
}

func Test_SameNumbers(t *testing.T) {
	in := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	res := threeSum(in)
	fmt.Println(len(res))
	for _, r := range res {
		for _, i := range r {
			fmt.Print(fmt.Sprintf("%v\t", i))
		}
		fmt.Print("\n")
	}
}
