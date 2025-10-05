// Copyright © 2025 Alex Temnok. All rights reserved.

package bitmap

type word = uint64

type Bitmap struct {
	width, height int
	words         []word
}

const (
	ones = ^word(0)
)

func New(w, h int) *Bitmap {
	w, h = max(w, 1), max(h, 1)

	b := &Bitmap{width: w, height: h}
	b.words = make([]word, b.addr(w, h-1)+1)
	return b
}

func (b *Bitmap) Clone() *Bitmap {
	return &Bitmap{
		width:  b.width,
		height: b.height,
		words:  append([]word{}, b.words...),
	}
}

func (b *Bitmap) Width() int {
	return b.width
}

func (b *Bitmap) Height() int {
	return b.height
}

func (b *Bitmap) Set1(x0, x1, y int) {
	b.Set(x0, x1, y, 1)
}

func (b *Bitmap) Set0(x0, x1, y int) {
	b.Set(x0, x1, y, 0)
}

func (b *Bitmap) Set(x0, x1, y, bit int) {
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

	if (bit & 1) != 0 {
		if i0 == i1 {
			b.words[i0] |= mask(j0, j1)

			return
		}

		b.words[i0] |= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.words[i] = ones
		}

		b.words[i1] |= mask(0, j1)
	} else {
		if i0 == i1 {
			b.words[i0] &^= mask(j0, j1)

			return
		}

		b.words[i0] &^= mask(j0, 64)
		for i := i0 + 1; i < i1; i++ {
			b.words[i] = 0
		}

		b.words[i1] &^= mask(0, j1)
	}
}

func (b *Bitmap) Get(x, y int) int {
	return int(b.words[b.addr(x, y)]>>(x%64)) & 1
}

func (b *Bitmap) Xor(a *Bitmap) {
	n := min(len(a.words), len(b.words))
	for i, w := range a.words[:n] {
		b.words[i] ^= w
	}
}

func (b *Bitmap) Or(a *Bitmap) {
	n := min(len(a.words), len(b.words))
	for i, w := range a.words[:n] {
		b.words[i] |= w
	}
}

func (b *Bitmap) addr(x, y int) int {
	return ((b.width+63)/64)*y + x/64
}

func mask(i, j int) word {
	return (ones << uint64(i)) ^ (ones << uint64(j))
}
