package bitmap

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/geom"
	"temnok/lab/path"
	"testing"
)

func savePng(t *testing.T, name string, im image.Image) {
	_ = os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
	assert.NoError(t, err)

	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}

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

func TestBitmap_SavePng(t *testing.T) {
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

	savePng(t, "chess.png", bm.ToImage(color.Black, color.White))
}

func TestBitmap_SaveBrush(t *testing.T) {
	b := NewBitmap(40, 40)
	b.Segments(20, 20, NewRoundBrush(20))

	savePng(t, "brush.png", b.ToImage(color.Black, color.White))
}

func TestBitmap_SaveBezier(t *testing.T) {
	b := NewBitmap(1000, 1000)
	brush := NewRoundBrush(20)

	b.Segments(350, 250, brush)
	b.Segments(650, 250, brush)
	path.Visit([]geom.XY{{250, 500}, {250, 750}, {750, 750}, {750, 500}}, VisitDotted(40, func(x, y int) {
		b.Segments(x, y, brush)
	}))

	bg := color.RGBA{R: 0, G: 0x80, B: 0, A: 0xff}
	fg := color.RGBA{R: 0xff, G: 0x80, B: 0, A: 0xff}
	savePng(t, "bezier.png", b.ToImage(bg, fg))
}

func TestBitmap_SaveRect(t *testing.T) {
	b := NewBitmap(2000, 2000)
	brush := NewRoundBrush(10)

	path.Visit([]geom.XY{
		{200, 200}, {200, 200}, {1800, 200},
		{1800, 200}, {1800, 200}, {1800, 1800},
		{1800, 1800}, {1800, 1800}, {200, 1800},
		{200, 1800}, {200, 1800}, {200, 200},
		{200, 200},
	}, func(x, y int) {
		b.Segments(x, y, brush)
	})

	path.Visit([]geom.XY{
		{500, 500}, {500, 500}, {1500, 500},
		{1500, 500}, {1500, 500}, {1500, 1500},
		{1500, 1500}, {1500, 1500}, {500, 1500},
		{500, 1500}, {500, 1500}, {500, 500},
		{500, 500},
	}, VisitDotted(20, func(x, y int) {
		b.Segments(x, y, brush)
	}))

	savePng(t, "rect.png", b.ToImage(color.Black, color.White))
}
