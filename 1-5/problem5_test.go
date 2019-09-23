package problem5

import (
	"testing"
)

func TestOneEditAway(t *testing.T) {
	cases := []struct {
		inputA   string
		inputB   string
		expected bool
	}{
		{inputA: "asdf", inputB: "asdf", expected: true},  // same
		{inputA: "asdf", inputB: "asgf", expected: true},  // one change
		{inputA: "asdf", inputB: "asf", expected: true},   // one add
		{inputA: "asf", inputB: "asdf", expected: true},   // one removeal
		{inputA: "asdf", inputB: "fdsa", expected: false}, // too far
		{inputA: "", inputB: "", expected: true},
		{inputA: " ", inputB: "", expected: true},
		{inputA: "", inputB: " ", expected: true},
	}

	for _, c := range cases {
		actual := oneEditAway(c.inputA, c.inputB)
		if actual != c.expected {
			t.Fatalf("InputA: %s, InputB: %s. Expected: %t, Actual: %t",
				c.inputA,
				c.inputB,
				c.expected,
				actual,
			)
		}
	}
}
