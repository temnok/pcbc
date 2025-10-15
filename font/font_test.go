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

func _TestConvertX(t *testing.T) {
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
