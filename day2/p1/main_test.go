package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restore(t *testing.T) {
	tests := []struct {
		desc     string
		codes    []int
		expected []int
	}{
		{
			desc:     "explanation example",
			codes:    []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			desc:     "first example",
			codes:    []int{1, 0, 0, 0, 99},
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			desc:     "second example",
			codes:    []int{2, 3, 0, 3, 99},
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			desc:     "third example",
			codes:    []int{2, 4, 4, 5, 99, 0},
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			desc:     "fourth example",
			codes:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := restore(test.codes)
			assert.Equal(t, test.expected, actual)
		})
	}
}
