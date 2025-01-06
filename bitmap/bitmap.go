// Copyright © 2025 Alex Temnok. All rights reserved.

package bitmap

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Bitmap struct {
	elems         []uint64
	width, height int
}

func NewBitmap(w, h int) *Bitmap {
	if w <= 0 || h <= 0 {
		panic(fmt.Errorf("invalid bitmap width=%v or height=%v", w, h))
	}

	b := &Bitmap{width: w, height: h}
	b.elems = make([]uint64, b.addr(w, h))
	return b
}

func (b *Bitmap) Invert() {
	for i := range b.elems {
		b.elems[i] = ^b.elems[i]
	}
}

func (b *Bitmap) Set1(x0, x1, y int) {
	b.Set(x0, x1, y, 1)
}

func (b *Bitmap) Set0(x0, x1, y int) {
	b.Set(x0, x1, y, 0)
}

func (b *Bitmap) Set(x0, x1, y, val int) {
	if x0 < 0 {
		x0 = 0
	}

	if x1 > b.width {
		x1 = b.width
	}

	if x0 >= x1 || y < 0 || y >= b.height {
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

func (b *Bitmap) Get(x, y int) int {
	return int(b.elems[b.addr(x, y)]>>(x%64)) & 1
}

//go:inline
func (b *Bitmap) addr(x, y int) int {
	return ((b.width+63)/64)*y + x/64
}

func (b *Bitmap) ToImage(zero, one color.Color) image.Image {
	return NewBitmapsImage([]*Bitmap{b}, [][2]color.Color{{zero, one}}, false)
}

//go:inline
func mask(i, j int) uint64 {
	return (math.MaxUint64 << uint64(i)) ^ (math.MaxUint64 << uint64(j))
}
