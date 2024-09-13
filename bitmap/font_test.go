package bitmap

import (
	"image/color"
	"temnok/lab/font"
	"temnok/lab/line"
	"temnok/lab/twod"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 200.0

	b := NewBitmap(16*scale*font.Width, 6*scale)
	brush := NewRoundBrush(font.Normal * scale)

	fontData := font.Lines

	transform := twod.Identity() // t2d.Transform{{1, 0}, {-0.25, 1}}

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			if fontData[c] == nil {
				continue
			}

			x0, y0 := float64(j)*scale*font.Width, float64(i)*scale

			for _, stroke := range fontData[c] {
				var px, py float64

				for step, p := range stroke {
					x := x0 + p.X*scale
					y := y0 + p.Y*scale

					v := transform.Point(twod.Coord{X: x, Y: y})
					x, y = v.X, v.Y

					if step != 0 {
						//bezier.CubicVisit([]twod.Coord{{px, py}, {px, py}, {x, y}, {x, y}}, func(x, y int) {
						//	b.Segments(x, y, brush)
						//})
						line.Visit(int(px), int(py), int(x), int(y), func(x, y int) {
							b.Segments(x, y, brush)
						})
					}
					px, py = x, y
				}
			}
		}
	}

	savePng(t, "font.png", b.ToImage(color.White, color.Black))
}
