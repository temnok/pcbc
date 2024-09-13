package line

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		x0, y0, x1, y1 int
		expected       [][2]int
	}{
		{x0: 0, y0: 0, x1: 0, y1: 0, expected: [][2]int{{0, 0}}},
		{x0: 2, y0: 0, x1: 0, y1: 0, expected: [][2]int{{2, 0}, {1, 0}, {0, 0}}},
		{x0: -1, y0: 1, x1: 1, y1: -1, expected: [][2]int{{-1, 1}, {0, 0}, {1, -1}}},
		{x0: 5, y0: 3, x1: 0, y1: 0, expected: [][2]int{{5, 3}, {4, 2}, {3, 2}, {2, 1}, {1, 1}, {0, 0}}},
	}

	for _, test := range tests {
		actual := [][2]int{}

		Visit(test.x0, test.y0, test.x1, test.y1, func(x, y int) {
			actual = append(actual, [2]int{x, y})
		})

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Visit(%v,%v,%v,%v):\nwant %v\n got %v\n",
				test.x0, test.y0, test.x1, test.y1, test.expected, actual)
		}
	}
}

func TestVisit_Random(t *testing.T) {
	for i := 0; i < 10_000; i++ {
		m := 1 + rand.IntN(1000)
		x0, y0, x1, y1 := rand.IntN(m), rand.IntN(m), rand.IntN(m), rand.IntN(m)

		expectedCalls := max(x0-x1, x1-x0, y0-y1, y1-y0) + 1
		actualCalls := 0

		Visit(x0, y0, x1, y1, func(x, y int) {
			actualCalls++
		})

		if actualCalls != expectedCalls {
			t.Errorf("Visit(%v,%v,%v,%v) calls:\nwant %v\n got %v\n",
				x0, y0, x1, y1, expectedCalls, actualCalls)
		}
	}
}
