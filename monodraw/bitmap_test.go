package monodraw

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/bezier"
	"testing"
)

func TestBitmap_Mask(t *testing.T) {
	tests := []struct {
		want uint64
		i, j int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{3, 0, 2},
		{7, 0, 3},
		{15, 0, 4},
		{0x18, 3, 5},
		{0x7FFF_FFFF_FFFF_FFFF, 0, 63},
		{0xFFFF_FFFF_FFFF_FFFE, 1, 64},
		{0xFFFF_FFFF_FFFF_FFFF, 0, 64},
		{0xFFFF_FFFF_FFFF_FFFF, 64, 0},
	}

	for _, test := range tests {
		got := mask(test.i, test.j)
		if test.want != got {
			t.Errorf("mask(%d, %d):\nwant %x\n got %x\n", test.i, test.j, test.want, got)
		}
	}
}

func xTestBitmap_SavePng(t *testing.T) {
	const d = 1_000
	bm := NewBitmap(8*d, 8*d)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				continue
			}

			x0, y0 := j*d, i*d
			for y := y0; y < y0+d; y++ {
				bm.Segment(x0, x0+d, y)
			}
		}
	}

	f, err := os.Create("chess.png")
	assert.NoError(t, err)

	im := bm.ToImage(color.Black, color.White)
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}

func xTestBitmap_SaveBrush(t *testing.T) {
	b := NewBitmap(40, 40)
	b.Segments(20, 20, NewRoundBrush(20))

	f, err := os.Create("brush.png")
	assert.NoError(t, err)

	im := b.ToImage(color.Black, color.White)
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}

func xTestBitmap_SaveBezier(t *testing.T) {
	b := NewBitmap(1000, 1000)
	brush := NewRoundBrush(20)

	b.Segments(350, 250, brush)
	b.Segments(650, 250, brush)
	bezier.CubicVisit([]bezier.Point{{250, 500}, {250, 750}, {750, 750}, {750, 500}}, VisitDotted(40, func(x, y int) {
		b.Segments(x, y, brush)
	}))

	f, err := os.Create("bezier.png")
	assert.NoError(t, err)

	im := b.ToImage(color.RGBA{0, 0x80, 0, 0xff}, color.RGBA{0xff, 0x80, 0, 0xff})
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}

func xTestBitmap_SaveRect(t *testing.T) {
	b := NewBitmap(2000, 2000)
	brush := NewRoundBrush(10)

	bezier.CubicVisit([]bezier.Point{
		{200, 200}, {200, 200}, {1800, 200},
		{1800, 200}, {1800, 200}, {1800, 1800},
		{1800, 1800}, {1800, 1800}, {200, 1800},
		{200, 1800}, {200, 1800}, {200, 200},
		{200, 200},
	}, func(x, y int) {
		b.Segments(x, y, brush)
	})

	bezier.CubicVisit([]bezier.Point{
		{500, 500}, {500, 500}, {1500, 500},
		{1500, 500}, {1500, 500}, {1500, 1500},
		{1500, 1500}, {1500, 1500}, {500, 1500},
		{500, 1500}, {500, 1500}, {500, 500},
		{500, 500},
	}, VisitDotted(20, func(x, y int) {
		b.Segments(x, y, brush)
	}))

	f, err := os.Create("rect.png")
	assert.NoError(t, err)

	im := b.ToImage(color.Black, color.White)
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}
