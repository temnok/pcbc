// Copyright © 2025 Alex Temnok. All rights reserved.

package cbc

import (
	"math"
	"reflect"
	"testing"
)

func TestRasterize(t *testing.T) {
	tests := []struct {
		xy   []float64
		want [][2]int
	}{
		{
			xy:   []float64{0, 0, 0, 0, 0, 0, 0, 0},
			want: nil,
		},
		{
			xy:   []float64{0, 0, 0, 0, 3, 3, 3, 3},
			want: [][2]int{{1, 1}, {2, 2}, {3, 3}},
		},
		{
			xy:   []float64{0, 0, 3, 0, 5, 3, 5, 5},
			want: [][2]int{{1, 0}, {2, 1}, {3, 1}, {4, 2}, {5, 3}, {5, 4}, {5, 5}},
		},
		{
			xy:   circle(0.5),
			want: [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}},
			/*
			  . 0 .
			  1 + 3
			  . 2 .
			*/
		},
		{
			xy:   circle(1),
			want: [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}},
			/*
			  . 0 .
			  1 + 3
			  . 2 .
			*/
		},
		{
			xy:   circle(1.5),
			want: [][2]int{{1, 1}, {0, 2}, {-1, 1}, {-2, 0}, {-1, -1}, {0, -2}, {1, -1}, {2, 0}},
			/*
			  . . 1 . .
			  . 2 . 0 .
			  3 . + . 7
			  . 4 . 6 .
			  . . 5 . .
			*/
		},
		{
			xy:   circle(2),
			want: [][2]int{{1, 1}, {0, 2}, {-1, 1}, {-2, 0}, {-1, -1}, {0, -2}, {1, -1}, {2, 0}},
			/*
			  . . 1 . .
			  . 2 . 0 .
			  3 . + . 7
			  . 4 . 6 .
			  . . 5 . .
			*/
		},
		{
			xy:   circle(2.5),
			want: [][2]int{{2, 1}, {2, 2}, {1, 2}, {0, 3}, {-1, 2}, {-2, 2}, {-2, 1}, {-3, 0}, {-2, -1}, {-2, -2}, {-1, -2}, {0, -3}, {1, -2}, {2, -2}, {2, -1}, {3, 0}},
			/*
			  . . . 3 . . .
			  . 5 4 . 2 1 .
			  . 6 . . . 0 .
			  7 . . + . . f
			  . 8 . . . e .
			  . 9 a . c d .
			  . . . b . . .
			*/
		},
		{
			xy:   circle(3),
			want: [][2]int{{3, 1}, {2, 2}, {1, 3}, {0, 3}, {-1, 3}, {-2, 2}, {-3, 1}, {-3, 0}, {-3, -1}, {-2, -2}, {-1, -3}, {0, -3}, {1, -3}, {2, -2}, {3, -1}, {3, 0}},
			/*
			  . . 4 3 2 . .
			  . 5 . . . 1 .
			  6 . . . . . 0
			  7 . . + . . f
			  8 . . . . . e
			  . 9 . . . d .
			  . . a b c . .
			*/
		},
	}

	for _, test := range tests {
		var got [][2]int

		for i := 0; i+8 <= len(test.xy); i += 6 {
			Rasterize(test.xy[i:i+8], func(x, y int) {
				got = append(got, [2]int{x, y})
			})
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Rasterize(%v):\nwant %v\n got %#v", test.xy, test.want, got)
		}
	}
}

func circle(r float64) []float64 {
	k := 4 * (math.Sqrt(2) - 1) / 3
	m := r * k

	return []float64{
		r, 0,
		r, m, m, r,
		0, r,
		-m, r, -r, m,
		-r, 0,
		-r, -m, -m, -r,
		0, -r,
		m, -r, r, -m,
		r, 0,
	}
}
