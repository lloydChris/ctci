package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abcba", true},
		{"abccba", true},
		{"asdf", false},
	}

	for _, c := range testCases {
		actual := isPalindrome(c.input)
		if actual != c.expected {
			t.Fatalf("Input: %s. Expected: %v. Actual: %v", c.input, c.expected, actual)
		}
	}
}
