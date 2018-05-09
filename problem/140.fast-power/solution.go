package lintcode

/**
 * @param a: A 32bit integer
 * @param b: A 32bit integer
 * @param n: A 32bit integer
 * @return: An integer
 */
func fastPower(a int, b int, n int) int {
	// (a*b)%n=((a%n)*(b%n))%n
	// a^n=(a^(n/2))*(a^(n/2))=...
	if n == 0 {
		return 1 % b
	}
	if n == 1 {
		return a % b
	}
	halfPower := fastPower(a, b, n/2)
	if n%2 == 0 {
		return (halfPower * halfPower) % b
	}
	return (((halfPower * halfPower) % b) * (a % b)) % b
}
