package bitmap

import "math"

type Segment struct {
	X0, X1, Y int16
}

/*

 *
***
 *


*/

func NewRoundBrush(r int) []Segment {
	if r <= 0 {
		return nil
	}

	brush := make([]Segment, r*2+1)

	rr := (r + 1) * r
	for y := 0; y <= r; y++ {
		x := 0
		for (x+1)*(x+1)+y*y <= rr {
			x++
		}

		brush[r-y] = Segment{X0: int16(-x), X1: int16(x + 1), Y: int16(-y)}
		brush[r+y] = Segment{X0: int16(-x), X1: int16(x + 1), Y: int16(y)}
	}

	return brush
}

func NewRoundBrushOld(ri int) []Segment {
	d := ri * 2
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
