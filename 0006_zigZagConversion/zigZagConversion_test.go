package zigZagConversion

import (
	"fmt"
	"testing"
)

func Test_Convert(t *testing.T) {
	/*
		P   A   H   N
		A P L S I I G
		Y   I   R

	*/

	input := "PAYPALISHIRING"
	numRows := 3
	output := "PAHNAPLSIIGYIR"

	fmt.Println(len(input), len([]byte(input)))

	res := convert(input, numRows)
	if res != output {
		t.Errorf("expected output '%s' but got '%s'", output, res)
	}
}

func Test_Convert4Rows(t *testing.T) {
	/*
		P     I    N
		A   L S  I G
		Y A   H R
		P     I

	*/

	input := "PAYPALISHIRING"
	numRows := 4
	output := "PINALSIGYAHRPI"

	res := convert(input, numRows)
	if res != output {
		t.Errorf("expected output '%s' but got '%s'", output, res)
	}
}

func Test_ConvertABC(t *testing.T) {
	/*
		a   e
		b d f
		c		g
	*/

	input := "abcdefg"
	numRows := 3
	output := "aebdfcg"

	res := convert(input, numRows)
	if res != output {
		t.Errorf("expected output '%s' but got '%s'", output, res)
	}
}

func Test_ConvertOneRow(t *testing.T) {
	input := "abcdefg"
	numRows := 1
	output := "abcdefg"

	res := convert(input, numRows)
	if res != output {
		t.Errorf("expected output '%s' but got '%s'", output, res)
	}
}
