package problem3

import (
	"testing"
)

func TestUrlify(t *testing.T) {
	testCases := make(map[string]string)
	testCases["asdf asdf asdf"] = "asdf%20asdf%20asdf"
	testCases["  asdf  "] = "%20%20asdf%20%20"

	for input, expected := range testCases {
		actual := urlify(input, len(input))
		if actual != expected {
			t.Error()
		}
	}
}
