package lintcode

var array [][]int

/**
 * @param n: An integer
 * @return: a square matrix
 */
func generateMatrix(n int) [][]int {
	// init the slice
	array = make([][]int, n)
	for i, _ := range array {
		array[i] = make([]int, n)
	}
	generateMatrixByNumber(n, 0, 1)
	return array
}

func generateMatrixByNumber(n int, row int, start int) {
	num := start
	if n == row*2+1 {
		array[row][row] = num
		return
	}
	for i := row; i <= n-1-row; i++ {
		array[row][i] = num
		num++
	}
	for i := row + 1; i <= n-1-row-1; i++ {
		array[i][n-1-row] = num
		num++
	}
	for i := n - 1 - row; i >= row; i-- {
		array[n-1-row][i] = num
		num++
	}
	for i := n - 1 - row - 1; i > row; i-- {
		array[i][row] = num
		num++
	}
	if num <= n*n {
		generateMatrixByNumber(n, row+1, num)
	}
}
