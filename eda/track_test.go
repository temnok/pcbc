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

		want Track
	}{
		{
			p0: path.Point{0, 0}, p1: path.Point{0, 0}, steps: []float64{1},
			want: Track{},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 0}, steps: []float64{},
			want: Track{{0, 0}, {1, 0}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 1}, steps: []float64{},
			want: Track{{0, 0}, {1, 1}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{-1, +1},
			want: Track{{0, 0}, {0, 1}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{},
			want: Track{{0, 0}, {1, 1}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{-1},
			want: Track{{0, 0}, {0, 1}, {1, 2}},
		},

		{
			p0: path.Point{0, 0}, p1: path.Point{1, 2}, steps: []float64{1},
			want: Track{{0, 0}, {1, 0}, {1, 2}},
		},

		{
			p0: path.Point{-6, 5}, p1: path.Point{-10, 6}, steps: []float64{-1},
			want: Track{{-6, 5}, {-7, 5}, {-9, 5}, {-10, 6}},
		},
	}

	for _, test := range tests {
		got := TrackV2(test.p0, test.p1, test.steps...)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("TrackV2(%v, %v, %v):\nwant %v\n got %v", test.p0, test.p1, test.steps, test.want, got)
		}
	}
}
