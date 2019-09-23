package problem4

func isPalindrome(input string) bool {
	mappedString := make(map[rune]int)

	for _, char := range input {
		if char != ' ' { //Ignore spaces
			mappedString[char]++
		}
	}

	var numberOfOddRunCounts int = 0

	for _, count := range mappedString {
		if count%2 != 0 {
			numberOfOddRunCounts++
		}
	}

	if numberOfOddRunCounts > 1 {
		return false
	}
	return true
}
