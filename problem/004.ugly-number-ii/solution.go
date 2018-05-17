package lintcode

var (
	uglyNumberArray []int
)

/**
 * @param n: An integer
 * @return: the nth prime number as description.
 */
func nthUglyNumber(n int) int {
	uglyNumberArray = []int{1}
	count2 := 0
	count3 := 0
	count5 := 0
	for len(uglyNumberArray) < n {
		num2 := uglyNumberArray[count2] * 2
		num3 := uglyNumberArray[count3] * 3
		num5 := uglyNumberArray[count5] * 5
		minNum := min(num2, num3, num5)
		if minNum == num2 {
			count2++
		} else if minNum == num3 {
			count3++
		} else if minNum == num5 {
			count5++
		}
		// ship the duplicate elements
		if minNum == uglyNumberArray[len(uglyNumberArray)-1] {
			continue
		} else {
			uglyNumberArray = append(uglyNumberArray, minNum)
		}
	}
	return uglyNumberArray[n-1]
}

func min(a int, b int, c int) int {
	tmp := a
	if b < tmp {
		tmp = b
	}
	if c < tmp {
		tmp = c
	}
	return tmp
}
