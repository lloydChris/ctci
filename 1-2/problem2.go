package problem2

import (
	"sort"
)

func buildStringMap(input string) map[rune]int {
	result := make(map[rune]int)
	for _, value := range input {
		currentCount, exists := result[value]
		if exists {
			result[value] = currentCount + 1
		} else {
			result[value] = 1
		}
	}
	return result
}

func checkPermutation(a string, b string) bool {
	// push each string into a map
	mapA := buildStringMap(a)
	mapB := buildStringMap(b)

	//Comapre string maps

	for key, value := range mapA {
		charCount, exist := mapB[key]
		if !exist {
			//missing char, not a permutation
			return false
		}
		if charCount != value {
			//wrong char count, not a permutation
			return false
		}
	}
	return true
}

func checkPermutationSort(a string, b string) bool {
	aRune := []rune(a)
	bRune := []rune(b)

	if len(aRune) != len(bRune) {
		return false
	}

	sort.Slice(aRune, func(i, j int) bool { return aRune[i] < aRune[j] })
	sort.Slice(bRune, func(i, j int) bool { return bRune[i] < bRune[j] })

	for i := range aRune {
		if aRune[i] != bRune[i] {
			return false
		}
	}
	return true
}
