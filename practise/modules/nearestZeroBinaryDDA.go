package modules

import "math"

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	INF := m + n
	for idx := 0; idx < m; idx++ {
		for jdx := 0; jdx < n; jdx++ {
			if mat[idx][jdx] == 0 {
				continue
			}

			top, left := INF, INF
			if idx-1 >= 0 {
				top = mat[idx-1][jdx]
			}
			if jdx-1 >= 0 {
				left = mat[idx][jdx-1]
			}
			mat[idx][jdx] = int(math.Min(float64(top), float64(left))) + 1
		}
	}

	for idx := m - 1; idx >= 0; idx-- {
		for jdx := n - 1; jdx >= 0; jdx-- {
			if mat[idx][jdx] == 0 {
				continue
			}

			bottom, right := INF, INF
			if idx+1 < m {
				bottom = mat[idx+1][jdx]
			}
			if jdx+1 < n {
				right = mat[idx][jdx+1]
			}
			mat[idx][jdx] = int(math.Min(float64(mat[idx][jdx]), math.Min(float64(bottom), float64(right))+1))
		}
	}

	return mat
}
