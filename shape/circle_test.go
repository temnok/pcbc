// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"reflect"
	"testing"
)

func TestCircle(t *testing.T) {
	tests := []struct {
		d    int
		want *Shape
	}{
		{
			d: 1,
			want: &Shape{
				rows: []row{
					{x0: 0, x1: 1, y: 0},
				},
			},
		},
		{
			d: 2,
			want: &Shape{
				rows: []row{
					{x0: -1, x1: 1, y: -1},
					{x0: -1, x1: 1, y: 0},
				},
			},
		},
		{
			/*
					 *
				   * 0 *
					 *
			*/
			d: 3,
			want: &Shape{
				rows: []row{
					{x0: 0, x1: 1, y: -1},
					{x0: -1, x1: 2, y: 0},
					{x0: 0, x1: 1, y: 1},
				},
			},
		},
		{
			/*
					 * *
				   * * * *
				   * * 0 *
					 * *
			*/
			d: 4,
			want: &Shape{
				rows: []row{
					{x0: -1, x1: 1, y: -2},
					{x0: -2, x1: 2, y: -1},
					{x0: -2, x1: 2, y: 0},
					{x0: -1, x1: 1, y: 1},
				},
			},
		},
		{
			/*
					 * * *
				   * * * * *
				   * * 0 * *
				   * * * * *
					 * * *
			*/
			d: 5,
			want: &Shape{
				rows: []row{
					{x0: -1, x1: 2, y: -2},
					{x0: -2, x1: 3, y: -1},
					{x0: -2, x1: 3, y: 0},
					{x0: -2, x1: 3, y: 1},
					{x0: -1, x1: 2, y: 2},
				},
			},
		},
		{
			/*
					 * * * *
				   * * * * * *
				   * * * * * *
				   * * * 0 * *
				   * * * * * *
					 * * * *
			*/
			d: 6,
			want: &Shape{
				rows: []row{
					{x0: -2, x1: 2, y: -3},
					{x0: -3, x1: 3, y: -2},
					{x0: -3, x1: 3, y: -1},
					{x0: -3, x1: 3, y: 0},
					{x0: -3, x1: 3, y: 1},
					{x0: -2, x1: 2, y: 2},
				},
			},
		},
		{
			/*
					   * * *
				     * * * * *
				   * * * * * * *
				   * * * 0 * * *
				   * * * * * * *
					 * * * * *
					   * * *
			*/
			d: 7,
			want: &Shape{
				rows: []row{
					{x0: -1, x1: 2, y: -3},
					{x0: -2, x1: 3, y: -2},
					{x0: -3, x1: 4, y: -1},
					{x0: -3, x1: 4, y: 0},
					{x0: -3, x1: 4, y: 1},
					{x0: -2, x1: 3, y: 2},
					{x0: -1, x1: 2, y: 3},
				},
			},
		},
		{
			/*
					     * * *
				     * * * * * * *
				     * * * * * * *
				   * * * * * * * * *
				   * * * * 0 * * * *
				   * * * * * * * * *
				     * * * * * * *
				     * * * * * * *
					     * * *
			*/
			d: 9,
			want: &Shape{
				rows: []row{
					{x0: -1, x1: 2, y: -4},
					{x0: -3, x1: 4, y: -3},
					{x0: -3, x1: 4, y: -2},
					{x0: -4, x1: 5, y: -1},
					{x0: -4, x1: 5, y: 0},
					{x0: -4, x1: 5, y: 1},
					{x0: -3, x1: 4, y: 2},
					{x0: -3, x1: 4, y: 3},
					{x0: -1, x1: 2, y: 4},
				},
			},
		},
	}

	for _, test := range tests {
		got := Circle(test.d)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Circle(%d):\nwant %+v\n got %+v", test.d, test.want, got)
		}
	}
}
