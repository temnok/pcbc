package monodraw

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

type Segment struct {
	X0, X1, Y int16
}

func NewBitmap(w, h int) *Bitmap {
	if w <= 0 || h <= 0 {
		panic(fmt.Errorf("invalid bitmap w=%v or h=%v", w, h))
	}

	return &Bitmap{
		elems: make([]uint64, ((w+63)/64)*h),
		w:     w,
		h:     h,
	}
}

func (b *Bitmap) Segment(x0, x1, y int) {
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

	if i0 == i1 {
		b.elems[i0] |= mask(x0%64, x1%64)

		return
	}

	b.elems[i0] |= mask(x0%64, 64)
	for i := i0 + 1; i < i1; i++ {
		b.elems[i] = math.MaxUint64
	}
	b.elems[i1] |= mask(0, (x1-1)%64+1)
}

func (b *Bitmap) Segments(x, y int, segs []Segment) {
	for _, seg := range segs {
		b.Segment(x+int(seg.X0), x+int(seg.X1), y+int(seg.Y))
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
