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
		expected []segment
	}{
		{
			contours: [][]point{
				{
					{0, 0},
				},
			},
			expected: []segment{
				{0, 0, 0},
				{0, 0, 0},
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
			expected: []segment{
				{0, 0, 0},
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 0},
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
			expected: []segment{
				{0, 1, 0},
				{0, 1, 1},
				{0, 1, 1},
				{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		got := []segment{}

		builder := New(func(x0, x1, y int) {
			got = append(got, segment{x0, x1, y})
		})
		for _, contour := range test.contours {
			for _, p := range contour {
				builder.AddPoint(p.x, p.y)
			}
			builder.FinishContour()
		}

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Builder(%v):\nwant %v\n got %v", test.contours, test.expected, got)
		}
	}
}
