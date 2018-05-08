package lintcode

var array []int

/**
 * @param n: An integer
 * @return: An array storing 1 to the largest number with n digits.
 */
func numbersByRecursion(n int) []int {
	recursion(n)
	return array[1:]
}

func recursion(n int) {
	if n == 0 {
		array = []int{0}
		return
	}
	recursion(n - 1)
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)
	array = make([]int, 0)
	for j := 0; j < len(arrayCopy); j++ {
		for i := 0; i < 10; i++ {
			array = append(array, 10*arrayCopy[j]+i)
		}
	}
}
