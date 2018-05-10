package lintcode

import (
	"reflect"
	"sort"
	"strings"
)

/**
 * @param A: a string
 * @param B: a string
 * @return: a boolean
 */
func Permutation(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	mapA := StringToMap(A)
	mapB := StringToMap(B)
	for k, v := range mapA {
		if v != mapB[k] {
			return false
		}
	}
	return true
}

func StringToMap(str string) map[byte]int {
	result := make(map[byte]int, len(str))
	for _, b := range str {
		result[byte(b)]++
	}
	return result
}

func Permutation2(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	bytesA := strings.Split(A, "")
	bytesB := strings.Split(B, "")
	sort.Strings(bytesA)
	sort.Strings(bytesB)
	return reflect.DeepEqual(bytesA, bytesB)
}

func Permutation3(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	var arrayA [256]int
	var arrayB [256]int
	for _, b := range A {
		arrayA[b]++
	}
	for _, b := range B {
		arrayB[b]++
	}
	return reflect.DeepEqual(arrayA, arrayB)
}
