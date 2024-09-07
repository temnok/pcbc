package monodraw

import "math"

type Segment struct {
	X0, X1, Y int16
}

func NewRoundBrush(d int) []Segment {
	if d <= 0 {
		return nil
	}

	res := make([]Segment, d)

	r := float64(d) / 2

	for i := 0; i < d; i++ {
		ir := float64(i) + 0.5 - r
		x := int(math.Round(math.Sqrt(float64(r*r - ir*ir))))
		res[i] = Segment{
			X0: int16(-x),
			X1: int16(x) + 1,
			Y:  int16(i - d/2),
		}
	}

	return res
}
