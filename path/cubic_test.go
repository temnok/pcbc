// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubicVisit_Random(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	rang := 1000.0

	for range 1000 {
		points := []Point{
			randomPoint(random, rang),
			randomPoint(random, rang),
			randomPoint(random, rang),
			randomPoint(random, rang),
		}

		p := points[0].Round()
		i := 0

		cubicVisit(points, func(x, y int) {
			c := Point{X: float64(x), Y: float64(y)}
			assert.NotEqual(t, p, c)

			if math.Abs(c.X-p.X) > 1 || math.Abs(c.Y-p.Y) > 1 {
				t.Fatalf("p=%#v,i=%v: %v -> %v", p, i, p, c)
			}

			p = c
			i++
		})

		assert.Equal(t, p, points[3].Round())
	}
}

func randomPoint(random *rand.Rand, rang float64) Point {
	return Point{
		X: random.Float64() * rang,
		Y: random.Float64() * rang,
	}
}
