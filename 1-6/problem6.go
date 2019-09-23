package problem6

import (
	"fmt"
	"strings"
)

func compressString(input string) string {
	var builder strings.Builder
	runningCount := 0
	inputLength := len(input)

	for i, char := range input {
		runningCount++
		if i+1 == inputLength { //last char, save and return
			fmt.Fprintf(&builder, "%c%d", char, runningCount)
		} else if rune(input[i+1]) != char {
			fmt.Fprintf(&builder, "%c%d", char, runningCount)
			runningCount = 0
		}
	}

	if len(input) < builder.Len() {
		return input
	}

	return builder.String()
}
