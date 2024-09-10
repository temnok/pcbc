package bezier

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubicVisit_Random(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	for range 10_000 {
		points := randomPoints(4, random, 2048)

		p := points[0].Round()
		i := 0

		CubicVisit(points, func(x, y int) {
			p1 := Point{float64(x), float64(y)}
			if i == 0 {
				assert.Equal(t, p, p1)
			} else {
				assert.NotEqual(t, p, p1)

				if math.Abs(p1.X-p.X) > 1 || math.Abs(p1.Y-p.Y) > 1 {
					t.Fatalf("p=%#v,i=%v: %v -> %v", p, i, p, p1)
				}
			}

			p = p1
			i++
		})

		assert.Equal(t, p, points[3].Round())
		assert.NotZero(t, i)
	}
}

func randomPoints(n int, random *rand.Rand, rang float64) []Point {
	points := make([]Point, n)
	for i := range points {
		points[i] = Point{
			random.Float64() * rang,
			random.Float64() * rang,
		}
	}
	return points
}
