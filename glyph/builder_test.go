package glyph

import (
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	type point = struct{ x, y int }
	type segment = struct{ x0, x1, y int }

	tests := []struct {
		contours [][]point
		expected *Glyph
	}{
		{
			contours: [][]point{
				{
					{0, 0},
				},
			},
			expected: &Glyph{
				y0: 0,
				y1: 1,
				data: []int16{
					// index
					2,
					4,
					// rows
					0, 0,
				},
			},
		},
		{
			contours: [][]point{
				{
					{0, 0},
					{0, 1},
					{0, 0},
				},
			},
			expected: &Glyph{
				y0: 0,
				y1: 2,
				data: []int16{
					// index
					3,
					5,
					7,
					// rows
					0, 0,
					0, 0,
				},
			},
		},
		{
			contours: [][]point{
				{
					{0, 0},
					{1, 0},
					{1, 1},
					{0, 1},
					{0, 0},
				},
			},
			expected: &Glyph{
				y0: 0,
				y1: 2,
				data: []int16{
					// index
					3,
					5,
					7,
					// rows
					0, 1,
					0, 1,
				},
			},
		},
		{
			contours: [][]point{
				{
					{0, 0},
					{1, 0},
					{2, 0},
					{2, 1},
					{2, 2},
					{1, 2},
					{0, 2},
					{0, 1},
					{0, 0},
				},
			},
			expected: &Glyph{
				y0: 0,
				y1: 3,
				data: []int16{
					// index
					4,
					6,
					8,
					10,
					// rows
					0, 2,
					0, 2,
					0, 2,
				},
			},
		},
	}

	for _, test := range tests {
		builder := New()
		for _, contour := range test.contours {
			for _, p := range contour {
				builder.AddPoint(p.x, p.y)
			}
			builder.FinishContour()
		}
		got := builder.Build()

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Builder(%+v):\nwant %+v\n got %+v", test.contours, test.expected, got)
		}
	}
}
