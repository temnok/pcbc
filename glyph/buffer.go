package glyph

import "sort"

type buffer struct {
	upper, lower [][]int
}

func (b *buffer) addSegment(x0, x1, y int) {
	rows := &b.upper
	if y < 0 {
		rows = &b.lower
		y = ^y
	}

	for len(*rows) <= y {
		*rows = append(*rows, nil)
	}
	(*rows)[y] = append((*rows)[y], x0, x1)
}

func (b *buffer) toGlyph() *Glyph {
	y0 := -len(b.lower)
	y1 := len(b.upper) - 1

	if y0 == 0 {
		for y0 < y1 && b.upper[y0] == nil {
			y0++
		}
	} else if y1 < 0 {
		for y0 < y1 && b.lower[^y1] == nil {
			y1--
		}
	}

	nRows := y1 - y0 + 1
	data := make([]int16, nRows+1)

	for y := y0; y <= y1; y++ {
		var row []int
		if y < 0 {
			row = b.lower[^y]
		} else {
			row = b.upper[y]
		}
		sort.Ints(row)

		data[y-y0] = int16(len(data))

		for i := 0; i+4 <= len(row); i += 4 {
			if i > 0 && row[i] == row[i-1] {
				data[len(data)-1] = int16(row[i+3])
				continue
			}
			data = append(data, int16(row[i]), int16(row[i+3]))
		}
	}

	data[nRows] = int16(len(data))
	return &Glyph{
		y0:   y0,
		y1:   y1 + 1,
		data: data,
	}
}
