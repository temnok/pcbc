package path

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestLinearVisit(t *testing.T) {
	tests := []struct {
		a, b     Point
		expected [][2]int
	}{
		{a: Point{0, 0}, b: Point{0, 0}, expected: [][2]int{}},
		{a: Point{2, 0}, b: Point{0, 0}, expected: [][2]int{{1, 0}, {0, 0}}},
		{a: Point{-1, 1}, b: Point{1, -1}, expected: [][2]int{{0, 0}, {1, -1}}},
		{a: Point{5, 3}, b: Point{0, 0}, expected: [][2]int{{4, 2}, {3, 2}, {2, 1}, {1, 1}, {0, 0}}},
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

		a := Point{float64(x0), float64(y0)}
		b := Point{float64(x1), float64(y1)}
		linearVisit(a, b, func(x, y int) {
			actualCalls++
		})

		if actualCalls != expectedCalls {
			t.Errorf("linearVisit(%v,%v) calls:\nwant %v\n got %v\n",
				a, b, expectedCalls, actualCalls)
		}
	}
}
