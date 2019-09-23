package problem2

import (
	"testing"
)

type testCases struct {
	first    string
	second   string
	expected bool
}

var cases = []testCases{
	testCases{
		first:    "this",
		second:   "that",
		expected: false,
	},
	testCases{
		first:    "qwer",
		second:   "rewq",
		expected: true,
	},
	testCases{
		first:    "qwerd",
		second:   "rewq",
		expected: false,
	},
}

func TestCheckPermutation(t *testing.T) {
	for _, c := range cases {
		actual := checkPermutation(c.first, c.second)
		if actual != c.expected {
			t.Error()
		}
	}
}

func TestCheckPermutationSort(t *testing.T) {
	for _, c := range cases {
		actual := checkPermutationSort(c.first, c.second)
		if actual != c.expected {
			t.Error()
		}
	}
}
