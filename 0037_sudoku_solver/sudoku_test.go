package sudoku

import (
	"encoding/json"
	"testing"
)

var inputBoard = `[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]`
var outputBoard = `[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]`


func Test_solveSudoku(t *testing.T) {
	input := parse(t, inputBoard)
	output := parse(t, outputBoard)

	solveSudoku(input)

	if !areSame(input, output) {
		t.Error("input and out put are not the same")
	}

}

func parse (t *testing.T, in string) [][]byte {
	var asStrings [][]string
	err := json.Unmarshal([]byte(in), &asStrings)
	if err != nil {
		t.Fatal(err)
	}

	out := make([][]byte, 0, len(asStrings))
	for _, row := range asStrings {
		byteRow := make([]byte, 0, len(row))
		for _, s := range row {
			byteRow = append(byteRow, []byte(s)[0])
		}
		out = append(out, byteRow)
	}

	return out
}


func areSame(l [][]byte, r [][]byte) bool {
	for rowID := range l {
		for colID := range l[rowID] {
			if 	l[rowID][colID] != r[rowID][colID] {
				return false
			}
		}
	}

	return true
}