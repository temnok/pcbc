package bitmap

type Segment struct {
	X0, X1, Y int16
}

func NewRoundBrush(d int) []Segment {
	if d <= 0 {
		return nil
	}

	brush := make([]Segment, d)

	off := (d + 1) % 2
	dd := (d - 1 - off) * (d - off)
	for y, r := 0, (d+1)/2; y < r; y++ {
		x := 0
		for ((x+1)*(x+1)+y*y)*4 <= dd {
			x++
		}

		brush[r-1-y] = Segment{X0: int16(-x - off), X1: int16(x + 1), Y: int16(-y - off)}
		brush[r-1+y+off] = Segment{X0: int16(-x - off), X1: int16(x + 1), Y: int16(y)}
	}

	return brush
}
