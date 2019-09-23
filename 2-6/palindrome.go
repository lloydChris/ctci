package palindrome

func recursiveIsPalindrome(input string, frontIndex int, backIndex int) bool {
	if frontIndex >= backIndex {
		return true
	}

	if !recursiveIsPalindrome(input, frontIndex+1, backIndex-1) {
		return false
	}

	return input[frontIndex] == input[backIndex]
}

func isPalindrome(input string) bool {
	startIndex := 0
	endIndex := len(input) - 1
	return recursiveIsPalindrome(input, startIndex, endIndex)
}

//aba
