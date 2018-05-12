package lintcode

/**
 * @param str: str: the given string
 * @return: char: the first unique character in a given string
 */
func firstUniqChar(str string) byte {
	strMap := make(map[rune]int, 0)
	for _, v := range str {
		strMap[v]++
	}
	for _, v := range str {
		if strMap[v] == 1 {
			return byte(v)
		}
	}
	return str[0]
}
