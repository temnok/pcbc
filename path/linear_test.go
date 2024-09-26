package path

import (
	"math/rand/v2"
	"reflect"
	"temnok/lab/geom"
	"testing"
)

func TestLinearVisit(t *testing.T) {
	tests := []struct {
		a, b     geom.XY
		expected [][2]int
	}{
		{a: geom.XY{0, 0}, b: geom.XY{0, 0}, expected: [][2]int{}},
		{a: geom.XY{2, 0}, b: geom.XY{0, 0}, expected: [][2]int{{1, 0}, {0, 0}}},
		{a: geom.XY{-1, 1}, b: geom.XY{1, -1}, expected: [][2]int{{0, 0}, {1, -1}}},
		{a: geom.XY{5, 3}, b: geom.XY{0, 0}, expected: [][2]int{{4, 2}, {3, 2}, {2, 1}, {1, 1}, {0, 0}}},
	}

	for _, test := range tests {
		actual := [][2]int{}

		linearVisit(test.a, test.b, func(x, y int) {
			actual = append(actual, [2]int{x, y})
		})

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("linearVisit(%v,%v):\nwant %v\n got %v\n",
				test.a, test.b, test.expected, actual)
		}
	}
}

func TestLinearVisit_Random(t *testing.T) {
	for i := 0; i < 10_000; i++ {
		m := 1 + rand.IntN(1000)
		x0, y0, x1, y1 := rand.IntN(m), rand.IntN(m), rand.IntN(m), rand.IntN(m)

		expectedCalls := max(x0-x1, x1-x0, y0-y1, y1-y0)
		actualCalls := 0

		a := geom.XY{float64(x0), float64(y0)}
		b := geom.XY{float64(x1), float64(y1)}
		linearVisit(a, b, func(x, y int) {
			actualCalls++
		})

		if actualCalls != expectedCalls {
			t.Errorf("linearVisit(%v,%v) calls:\nwant %v\n got %v\n",
				a, b, expectedCalls, actualCalls)
		}
	}
}
