package project7

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			input: [][]byte{
				{'a', 'b', 'c', 'd'},
				{'e', 'f', 'g', 'h'},
				{'i', 'j', 'k', 'l'},
				{'m', 'n', 'o', 'p'}},
			expected: [][]byte{
				{'m', 'i', 'e', 'a'},
				{'n', 'j', 'f', 'b'},
				{'o', 'k', 'g', 'c'},
				{'p', 'l', 'h', 'd'}},
		},
		{
			input: [][]byte{
				{'a', 'b', 'c'},
				{'e', 'f', 'g'},
				{'i', 'j', 'k'}},
			expected: [][]byte{
				{'i', 'e', 'a'},
				{'j', 'f', 'b'},
				{'k', 'g', 'c'}},
		},
	}

	for _, c := range cases {
		rotate(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: %v, Actual: %v\n", c.expected, c.input)
		}
	}

}
