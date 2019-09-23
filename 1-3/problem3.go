package problem3

import "strings"

func urlify(input string, length int) string {
	var space rune = 32
	var spaceReplacer = "%20"

	var result strings.Builder
	for _, char := range input {
		if char == space {
			result.WriteString(spaceReplacer)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
