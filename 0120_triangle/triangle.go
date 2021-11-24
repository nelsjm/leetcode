package triangle


func minimumTotal(triangle [][]int) int {
	for currentRow := len(triangle) - 1; currentRow >= 0; currentRow = currentRow - 1 {
		if currentRow == len(triangle) - 1 {
			continue
		}

		for i, v := range triangle[currentRow] {
			triangle[currentRow][i] = v + min(triangle[currentRow+1][i], triangle[currentRow+1][i+1])
		}
	}

	return triangle[0][0]
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}