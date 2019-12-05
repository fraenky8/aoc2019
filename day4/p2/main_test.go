package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbers_isInLargerGroup(t *testing.T) {
	tests := []struct {
		input    numbers
		expected bool
	}{
		{
			input:    []int{1, 1, 1, 1, 1, 1},
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 4, 4},
			expected: true,
		},
		{
			input:    []int{1, 1, 1, 1, 2, 2},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 2, 2, 2},
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 2, 2, 2},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 2, 2, 2, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 5},
			expected: false,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%v:%v", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := test.input.isInLargerGroup()
			assert.Equal(t, test.expected, actual)
		})
	}
}
