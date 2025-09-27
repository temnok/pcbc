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

		want []path.Point
	}{
		{
			p0: path.Point{0, -1}, p1: path.Point{-10, -5}, steps: []float64{},
			want: []path.Point{{0, -1}, {-4, -5}, {-10, -5}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{0, 0}, steps: []float64{1},
			want: nil,
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 0}, steps: []float64{},
			want: []path.Point{{0, 0}, {1, 0}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 1}, steps: []float64{},
			want: []path.Point{{0, 0}, {1, 1}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 1}, steps: []float64{1},
			want: []path.Point{{0, 0}, {1, 1}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{},
			want: []path.Point{{0, 0}, {1, 1}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{1},
			want: []path.Point{{0, 0}, {1, 1}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{0, 1},
			want: []path.Point{{0, 0}, {1, 0}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{2, 2}, steps: []float64{0, 1, -1, -1},
			want: []path.Point{{0, 0}, {1, 0}, {2, 1}, {2, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 0}, steps: []float64{-1},
			want: []path.Point{{0, 0}, {-1, 0}, {1, 0}},
		},
	}

	for _, test := range tests {
		got := LinearTrack(test.p0, test.p1, test.steps...)
		want := path.Linear(test.want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("LinearTrack(%v, %v, %v):\nwant %v\n got %v", test.p0, test.p1, test.steps, want, got)
		}
	}
}
