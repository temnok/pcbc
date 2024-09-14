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
				y0: 0,
				upper: []row{
					{x0: 0, x1: 1},
				},
			},
		},
		{
			d: 2,
			want: &Shape{
				y0: -1,
				upper: []row{
					{x0: -1, x1: 1},
					{x0: -1, x1: 1},
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
				y0: 0,
				lower: []row{
					{x0: 0, x1: 1},
				},
				upper: []row{
					{x0: -1, x1: 2},
					{x0: 0, x1: 1},
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
				y0: -1,
				lower: []row{
					{x0: -1, x1: 1},
				},
				upper: []row{
					{x0: -2, x1: 2},
					{x0: -2, x1: 2},
					{x0: -1, x1: 1},
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
				y0: 0,
				lower: []row{
					{x0: -2, x1: 3},
					{x0: -1, x1: 2},
				},
				upper: []row{
					{x0: -2, x1: 3},
					{x0: -2, x1: 3},
					{x0: -1, x1: 2},
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
				y0: -1,
				lower: []row{
					{x0: -3, x1: 3},
					{x0: -2, x1: 2},
				},
				upper: []row{
					{x0: -3, x1: 3},
					{x0: -3, x1: 3},
					{x0: -3, x1: 3},
					{x0: -2, x1: 2},
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
				y0: 0,
				lower: []row{
					{x0: -3, x1: 4},
					{x0: -2, x1: 3},
					{x0: -1, x1: 2},
				},
				upper: []row{
					{x0: -3, x1: 4},
					{x0: -3, x1: 4},
					{x0: -2, x1: 3},
					{x0: -1, x1: 2},
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
				y0: 0,
				lower: []row{
					{x0: -4, x1: 5},
					{x0: -3, x1: 4},
					{x0: -3, x1: 4},
					{x0: -1, x1: 2},
				},
				upper: []row{
					{x0: -4, x1: 5},
					{x0: -4, x1: 5},
					{x0: -3, x1: 4},
					{x0: -3, x1: 4},
					{x0: -1, x1: 2},
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
