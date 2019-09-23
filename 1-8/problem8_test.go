package problem8

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{1, 2, 3},
			},
			expected: [][]int{
				{1, 0, 3},
				{0, 0, 0},
				{1, 0, 3},
			},
		},
		{
			input: [][]int{
				{0, 2, 3},
				{4, 0, 6},
				{1, 2, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 1, 6},
				{1, 2, 1},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 1, 6},
				{1, 2, 1},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{0, 1, 6},
				{1, 2, 1},
			},
			expected: [][]int{
				{0, 2, 3},
				{0, 0, 0},
				{0, 2, 1},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 1, 6},
				{1, 2, 0},
			},
			expected: [][]int{
				{1, 2, 0},
				{4, 1, 0},
				{0, 0, 0},
			},
		},
	}

	for _, c := range cases {
		zero(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: %v, Actual: %v\n", c.expected, c.input)
		}
	}
}
