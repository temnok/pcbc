// Copyright © 2025 Alex Temnok. All rights reserved.

package bitmap

type Bitmap struct {
	width, height int
	bits          []uint64
}

const (
	ones = uint64((1 << 64) - 1)
)

func New(w, h int) *Bitmap {
	w, h = max(w, 1), max(h, 1)

	b := &Bitmap{width: w, height: h}
	b.bits = make([]uint64, b.addr(w, h))
	return b
}

func (b *Bitmap) Width() int {
	return b.width
}

func (b *Bitmap) Height() int {
	return b.height
}

func (b *Bitmap) Invert() {
	for i := range b.bits {
		b.bits[i] = ^b.bits[i]
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
			b.bits[i0] |= mask(j0, j1)

			return
		}

		b.bits[i0] |= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.bits[i] = ones
		}

		b.bits[i1] |= mask(0, j1)
	} else {
		if i0 == i1 {
			b.bits[i0] &^= mask(j0, j1)

			return
		}

		b.bits[i0] &^= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.bits[i] = 0
		}

		b.bits[i1] &^= mask(0, j1)
	}
}

func (b *Bitmap) Get(x, y int) int {
	return int(b.bits[b.addr(x, y)]>>(x%64)) & 1
}

func (b *Bitmap) addr(x, y int) int {
	return ((b.width+63)/64)*y + x/64
}

func mask(i, j int) uint64 {
	return (ones << uint64(i)) ^ (ones << uint64(j))
}
