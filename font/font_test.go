// Copyright © 2025 Alex Temnok. All rights reserved.

package font

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/bitmap"
	"github.com/temnok/pcbc/bitmap/image"
	"github.com/temnok/pcbc/shape"
	"github.com/temnok/pcbc/transform"
	"github.com/temnok/pcbc/util"
	"image/color"
	"sort"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 100.0

	const height = 1.0
	bm := bitmap.New(16*scale*Width, 20*scale*height)

	lightBrush := shape.Circle(Light * scale)

	d := Medium * scale
	mediumBrush := shape.Circle(int(d))

	boldBrush := shape.Circle(Bold * scale)

	for i := 0; i < 6; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			tf := transform.Move(Width/2+float64(j)*Width, height/2+float64(19-i)*height).ScaleUniformly(scale)
			lightBrush.ForEachPathsPixel(symbolPaths[c], tf, bm.Set1)

			tf = transform.Move(Width/2+float64(j)*Width, height/2+float64(12-i)*height).ScaleUniformly(scale)
			mediumBrush.ForEachPathsPixel(symbolPaths[c], tf, bm.Set1)

			tf = transform.Move(Width/2+float64(j)*Width, height/2+float64(5-i)*height).ScaleUniformly(scale)
			boldBrush.ForEachPathsPixel(symbolPaths[c], tf, bm.Set1)
		}
	}

	assert.NoError(t, util.SavePNG("out/font.png", image.NewSingle(bm, color.Black, color.White)))
}

func xTestConvertX(t *testing.T) {
	for i := 0x20; i < 0x7f; i++ {
		strokes := [][]byte{{}}
		//for _, j := range data[i] {
		//	if j < 0 {
		//		j = -j
		//		strokes = append(strokes, nil)
		//	}
		//
		//	n := len(strokes) - 1
		//	strokes[n] = append(strokes[n], j+11)
		//}
		sort.Slice(strokes, func(i, j int) bool {
			return strokes[i][0] < strokes[j][0]
		})

		fmt.Printf("'%c': {", i)
		for i, s := range strokes {
			n := len(s)
			if n == 0 {
				continue
			}
			if s[0] > s[n-1] {
				for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
					s[l], s[r] = s[r], s[l]
				}
			}
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print("{")
			for j, b := range s {
				if j > 0 {
					fmt.Print(", ")
				}
				if b < 0 {
					b = -b
				}
				fmt.Print(b)
			}
			fmt.Print("}")
		}
		fmt.Println("},")
	}
}

func xTestConvertGrid(t *testing.T) {
	for ch, dat := range data {
		if len(dat) == 0 && ch != ' ' {
			continue
		}

		if ch%0x10 == 0 {
			fmt.Printf("\n// 0x%x\n", ch)
		}

		fmt.Printf("'%v': {", string(rune(ch)))

		for i, stroke := range dat {
			if i != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("{")

			for j, p := range stroke {
				if j != 0 {
					fmt.Printf(", ")
				}
				r, c := 10-p%10, p/10
				fmt.Printf("%v%v", r, c)
			}

			fmt.Printf("}")
		}

		fmt.Printf("},\n")
	}
}

func xTestCanonizeGrid(t *testing.T) {
	for ch, dat := range data {
		if len(dat) == 0 && ch != ' ' {
			continue
		}

		if ch%0x10 == 0 {
			fmt.Printf("\n// 0x%x\n", ch)
		}

		esc := ""
		if ch == '\'' || ch == '\\' {
			esc = "\\"
		}

		fmt.Printf("'%v%v': {", esc, string(rune(ch)))

		canonizeStrokes(dat)

		for i, stroke := range dat {
			if i != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("{")

			for j, p := range stroke {
				if j != 0 {
					fmt.Printf(", ")
				}
				r, c := p/10, p%10
				fmt.Printf("%v%v", r, c)
			}

			fmt.Printf("}")
		}

		fmt.Printf("},\n")
	}
}

func canonizeStrokes(strokes [][]byte) {
	for i, s := range strokes {
		strokes[i] = canonizeStroke(s)
	}

	sort.Slice(strokes, func(i, j int) bool {
		return strokeLess(strokes[i], strokes[j])
	})
}

func canonizeStroke(stroke []byte) []byte {
	n := len(stroke)

	if n == 1 {
		return stroke
	}

	if stroke[0] != stroke[n-1] {
		return strokeMin(stroke, strokeRevert(stroke))
	}

	stroke = stroke[:n-1]

	best := stroke
	for i := 1; i < len(stroke); i++ {
		if s := strokeShift(stroke, i); strokeLess(s, best) {
			best = s
		}
	}

	stroke = strokeRevert(stroke)
	for i := 0; i < len(stroke); i++ {
		if s := strokeShift(stroke, i); strokeLess(s, best) {
			best = s
		}
	}

	return append(best, best[0])
}

func strokeLess(a, b []byte) bool {
	for i, x := range a {
		if y := b[i]; x < y {
			return true
		} else if x > y {
			return false
		}
	}

	return len(a) < len(b)
}

func strokeMin(a, b []byte) []byte {
	if strokeLess(a, b) {
		return a
	} else {
		return b
	}
}

func strokeRevert(stroke []byte) []byte {
	s := append([]byte{}, stroke...)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func strokeShift(stroke []byte, shift int) []byte {
	s := make([]byte, len(stroke))
	for i := range stroke {
		s[i] = stroke[(i+shift)%len(stroke)]
	}

	return s
}
