package bitmap

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Bitmap struct {
	elems []uint64
	w, h  int
}

func NewBitmap(w, h int) *Bitmap {
	if w <= 0 || h <= 0 {
		panic(fmt.Errorf("invalid bitmap w=%v or h=%v", w, h))
	}

	b := &Bitmap{w: w, h: h}
	b.elems = make([]uint64, b.addr(w, h))
	return b
}

func (b *Bitmap) SetRow1(x0, x1, y int) {
	b.SetRowVal(x0, x1, y, 1)
}

func (b *Bitmap) SetRow0(x0, x1, y int) {
	b.SetRowVal(x0, x1, y, 0)
}

func (b *Bitmap) SetRowVal(x0, x1, y, val int) {
	if x0 < 0 {
		x0 = 0
	}

	if x1 > b.w {
		x1 = b.w
	}

	if x0 >= x1 || y < 0 || y >= b.h {
		return
	}

	i0, i1 := b.addr(x0, y), b.addr(x1-1, y)
	j0, j1 := x0%64, (x1-1)%64+1

	if (val & 1) != 0 {
		if i0 == i1 {
			b.elems[i0] |= mask(j0, j1)

			return
		}

		b.elems[i0] |= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.elems[i] = math.MaxUint64
		}
		b.elems[i1] |= mask(0, j1)
	} else {
		if i0 == i1 {
			b.elems[i0] &^= mask(j0, j1)

			return
		}

		b.elems[i0] &^= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.elems[i] = 0
		}
		b.elems[i1] &^= mask(0, j1)
	}
}

//go:inline
func (b *Bitmap) addr(x, y int) int {
	return ((b.w+63)/64)*y + x/64
}

func (b *Bitmap) ToImage(zero, one color.Color) image.Image {
	return &bitmapImage{
		bitmap:  b,
		palette: color.Palette{zero, one},
	}
}

//go:inline
func mask(i, j int) uint64 {
	return (math.MaxUint64 << uint64(i)) ^ (math.MaxUint64 << uint64(j))
}
