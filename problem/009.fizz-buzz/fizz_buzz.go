package lintcode

import (
	"strconv"
)

/**
 * @param n: An integer
 * @return: A list of strings.
 */
func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "fizz buzz")
		} else if i%3 == 0 {
			result = append(result, "fizz")
		} else if i%5 == 0 {
			result = append(result, "buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
