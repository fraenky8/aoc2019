package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		desc     string
		wire1    Wire
		wire2    Wire
		distance int
	}{
		{
			desc: "explanation example",
			wire1: Wire{
				Directions: []Direction{
					{right, 8},
					{up, 5},
					{left, 5},
					{down, 3},
				},
				Positions: make(map[Position]struct{}),
			},
			wire2: Wire{
				Directions: []Direction{
					{up, 7},
					{right, 6},
					{down, 4},
					{left, 4},
				},
				Positions: make(map[Position]struct{}),
			},
			distance: 6,
		},
		{
			desc: "first example",
			wire1: Wire{
				Directions: []Direction{
					{right, 75},
					{down, 30},
					{right, 83},
					{up, 83},
					{left, 12},
					{down, 49},
					{right, 71},
					{up, 7},
					{left, 72},
				},
				Positions: make(map[Position]struct{}),
			},
			wire2: Wire{
				Directions: []Direction{
					{up, 62},
					{right, 66},
					{up, 55},
					{right, 34},
					{down, 71},
					{right, 55},
					{down, 58},
					{right, 83},
				},
				Positions: make(map[Position]struct{}),
			},
			distance: 159,
		},
		{
			desc: "second example",
			wire1: Wire{
				Directions: []Direction{
					{right, 98},
					{up, 47},
					{right, 26},
					{down, 63},
					{right, 33},
					{up, 87},
					{left, 62},
					{down, 20},
					{right, 33},
					{up, 53},
					{right, 51},
				},
				Positions: make(map[Position]struct{}),
			},
			wire2: Wire{
				Directions: []Direction{
					{up, 98},
					{right, 91},
					{down, 20},
					{right, 16},
					{down, 67},
					{right, 40},
					{up, 7},
					{right, 15},
					{up, 6},
					{right, 7},
				},
				Positions: make(map[Position]struct{}),
			},
			distance: 135,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := intersect(test.wire1, test.wire2)
			assert.Equal(t, test.distance, actual.distance())
		})
	}
}
