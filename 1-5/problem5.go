package problem5

import "math"

func oneEditAway(inputA string, inputB string) bool {
	typeOfEdit := len(inputA) - len(inputB)

	if math.Abs(float64(typeOfEdit)) > 1 {
		// more than on edit away
		return false
	}

	lenA := len(inputA)
	lenB := len(inputB)

	foundDifference := false
	for i, char := range inputA {
		if (i == lenA-1 || i == lenB-1) && lenA != lenB { //last char to check
			if foundDifference {
				return false
			}
			return true
		}
		if inputA[i] != inputB[i] {
			if foundDifference { //More than one difference found
				return false
			}
			foundDifference = true
			if typeOfEdit == 1 {
				inputB = inputB[:i] + string(char) + inputB[i:]
			} else if typeOfEdit == -1 {
				inputB = inputB[:i] + inputB[i+1:]
			}
			//do nothing in the case of one "replace" away
		}
	}
	return true
}
