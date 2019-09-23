package problem6

import (
	"testing"
)

func TestCompressString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "abbbccccaa", expected: "a1b3c4a2"},
		{input: "abcd", expected: "abcd"},
	}

	for _, c := range cases {
		actual := compressString(c.input)
		if actual != c.expected {
			t.Fatalf("Input: %s. Expected: %s, Actual: %s", c.input, c.expected, actual)
		}
	}
}
