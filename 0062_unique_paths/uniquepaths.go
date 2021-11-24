package uniquepaths

func uniquePaths(m int, n int) int {
	var cache = initCache(m,n)

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if x == 0 && y == 0 {
				cache[y][x] = 1
				continue
			}

			// look left
			leftPaths := 0
			if 0 <= x - 1 {
				leftPaths = cache[y][x-1]
			}

			// look up
			upPaths := 0
			if 0 <= y - 1 {
				upPaths = cache[y-1][x]
			}
			// set cache to left paths & up paths
			cache[y][x] = leftPaths + upPaths
		}
	}


	return cache[m-1][n-1]
}

func initCache(m int, n int) [][]int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	return cache
}