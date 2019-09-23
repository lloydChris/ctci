package problem4

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "taco cat", expected: true},
		{input: "tcao  cat", expected: true},
	}

	for _, c := range cases {
		actual := isPalindrome(c.input)
		if actual != c.expected {
			t.Fatalf("Input: %s. Expected: %t, Actual: %t\n",
				c.input,
				c.expected,
				actual)
		}
	}

}
