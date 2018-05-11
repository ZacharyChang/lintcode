package lintcode

/**
 * @param matrix: a matrix of m x n elements
 * @return: an integer list
 */
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)

	m := len(matrix)
	if m == 0 {
		return result
	}
	// single line
	if m == 1 {
		result = append(result, matrix[0]...)
		return result
	}

	n := len(matrix[0])

	// single row
	if n == 1 {
		for i := 0; i < m; i++ {
			result = append(result, matrix[i][0])
		}
		return result
	}

	// multi lines and rows
	for i := 0; i < n; i++ { // top
		result = append(result, matrix[0][i])
	}
	for i := 1; i < m; i++ { // right
		result = append(result, matrix[i][n-1])
	}
	for i := n - 2; i >= 0; i-- { // bottom
		result = append(result, matrix[m-1][i])
	}
	for i := m - 2; i > 0; i-- { // left
		result = append(result, matrix[i][0])
	}

	// recursive
	if m-2 >= 1 && n-2 >= 1 {
		nextMatrix := make([][]int, 0)
		for i := 1; i < m-1; i++ {
			line := make([]int, 0)
			for j := 1; j < n-1; j++ {
				line = append(line, matrix[i][j])
			}
			nextMatrix = append(nextMatrix, line)
		}
		result = append(result, spiralOrder(nextMatrix)...)
	}
	return result
}
