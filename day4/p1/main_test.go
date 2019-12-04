package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_digits(t *testing.T) {
	tests := []struct {
		desc     string
		input    int
		expected numbers
	}{
		{
			desc:     "",
			input:    123456,
			expected: numbers{1, 2, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := digits(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNumbers_isIncreasing(t *testing.T) {
	tests := []struct {
		input    numbers
		expected bool
	}{
		{
			input:    []int{1, 1, 1, 1, 1, 1},
			expected: true,
		},
		{
			input:    []int{2, 2, 3, 4, 5, 0},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 4, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 3, 6},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 4, 2},
			expected: false,
		},
		{
			input:    []int{1, 0, 3, 4, 4, 4},
			expected: false,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%v:%v", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := test.input.isIncreasing()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNumbers_hasDouble(t *testing.T) {
	tests := []struct {
		input    numbers
		expected bool
	}{
		{
			input:    []int{1, 1, 1, 1, 1, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 7, 8, 9},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: false,
		},
		{
			input:    []int{1, 1, 3, 4, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 2, 4, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 3, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 4, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 5},
			expected: true,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%v:%v", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := test.input.hasDouble()
			assert.Equal(t, test.expected, actual)
		})
	}
}
