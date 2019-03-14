package zigZagConversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	if len(s) <= numRows {
		return s
	}

	sourceBits := []byte(s)
	rows := make([][]byte, numRows)

	currRow := 0
	direction := 1
	for _, b := range sourceBits {
		r := rows[currRow]
		r = append(r, b)

		rows[currRow] = r

		if currRow == 0 {
			direction = 1
		}

		if currRow == numRows-1 {
			direction = -1
		}

		currRow += direction
	}

	destBits := make([]byte, 0, len(s))
	for _, r := range rows {
		destBits = append(destBits, r...)
	}

	return string(destBits)
}
