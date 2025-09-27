// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"reflect"
	"temnok/pcbc/path"
	"testing"
)

func TestTrack(t *testing.T) {
	tests := []struct {
		p0, p1 path.Point
		steps  []float64

		want path.Path
	}{
		{
			p0: path.Point{0, -1}, p1: path.Point{-10, -5}, steps: []float64{},
			want: path.Path{
				{0, -1}, {0, -1},
				{-4, -5}, {-4, -5}, {-4, -5},
				{-10, -5}, {-10, -5},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{0, 0}, steps: []float64{1},
			want: path.Path{},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 0}, steps: []float64{},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 0}, {1, 0},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 1}, steps: []float64{},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 1}, {1, 1},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 1}, steps: []float64{1},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 1}, {1, 1},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 1}, {1, 1}, {1, 1},
				{1, 2}, {1, 2},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{1},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 1}, {1, 1}, {1, 1},
				{1, 2}, {1, 2},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{0, 1},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 0}, {1, 0}, {1, 0},
				{1, 2}, {1, 2},
			},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{2, 2}, steps: []float64{0, 1, -1, -1},
			want: path.Path{
				{0, 0}, {0, 0},
				{1, 0}, {1, 0}, {1, 0},
				{2, 1}, {2, 1}, {2, 1},
				{2, 2}, {2, 2},
			},
		},
	}

	for _, test := range tests {
		got := Track(test.p0, test.p1, test.steps...)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Track(%v, %v, %v):\nwant %v\n got %v", test.p0, test.p1, test.steps, test.want, got)
		}
	}
}
