// Copyright © 2025 Alex Temnok. All rights reserved.

package font

import (
	"math"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type Align float64

const (
	Width = 0.6 // relative to height 1.0

	Light  = 0.08
	Medium = 0.1
	Bold   = 0.13

	AlignLeft   Align = 0.0
	AlignCenter Align = 0.5
	AlignRight  Align = 1.0
)

/*
Grid:
	19 29 39 49 59
	18 28 38 48 58
	17 27 37 47 57
	16 26 36 46 56
	15 25 35 45 55
	14 24 34 44 54
	13 23 33 43 53
	12 22 32 42 52
	11 21 31 41 51
*/

var data = [][][]byte{
	// 0x20
	' ':  {},
	'!':  {{33}, {35, 39}},
	'"':  {{27, 29}, {47, 49}},
	'#':  {{15, 55}, {17, 57}, {23, 29}, {43, 49}},
	'$':  {{13, 22, 42, 53, 54, 45, 25, 16, 17, 28, 48, 57}, {31, 39}},
	'%':  {{13, 59}, {19, 29, 28, 18, 19}, {44, 54, 53, 43, 44}},
	'&':  {{53, 18, 29, 39, 48, 15, 14, 23, 33, 55}},
	'\'': {{37, 39}},
	'(':  {{33, 25, 27, 39}},
	')':  {{33, 45, 47, 39}},
	'*':  {{15, 57}, {17, 55}, {34, 38}},
	'+':  {{16, 56}, {34, 38}},
	',':  {{21, 33}},
	'-':  {{16, 56}},
	'.':  {{33}},
	'/':  {{13, 59}},

	// 0x30
	'0': {{14, 18, 29, 49, 58, 54, 43, 23, 14}, {15, 25, 47, 57}},
	'1': {{13, 53}, {17, 39, 33}},
	'2': {{18, 29, 49, 58, 57, 13, 53}},
	'3': {{14, 23, 43, 54, 55, 46, 36, 46, 57, 58, 49, 29, 18}},
	'4': {{39, 15, 55}, {53, 57}},
	'5': {{13, 43, 54, 55, 46, 16, 19, 59}},
	'6': {{16, 46, 55, 54, 43, 23, 14, 16, 39}},
	'7': {{19, 59, 23}},
	'8': {{26, 17, 18, 29, 49, 58, 57, 46, 26, 15, 14, 23, 43, 54, 55, 46}},
	'9': {{33, 56, 58, 49, 29, 18, 17, 26, 56}},
	':': {{33}, {38}},
	';': {{21, 33}, {38}},
	'<': {{53, 16, 59}},
	'=': {{15, 55}, {17, 57}},
	'>': {{13, 56, 19}},
	'?': {{18, 29, 49, 58, 57, 35}, {33}},

	// 0x40
	'@':  {{43, 23, 14, 18, 29, 49, 58, 55, 35, 37, 57}},
	'A':  {{13, 18, 29, 49, 58, 53}, {16, 56}},
	'B':  {{16, 46, 57, 58, 49, 19, 13, 43, 54, 55, 46}},
	'C':  {{54, 43, 23, 14, 18, 29, 49, 58}},
	'D':  {{13, 19, 49, 58, 54, 43, 13}},
	'E':  {{16, 46}, {53, 13, 19, 59}},
	'F':  {{13, 19, 59}, {16, 46}},
	'G':  {{36, 56, 54, 43, 23, 14, 18, 29, 39, 49, 58}},
	'H':  {{13, 19}, {16, 56}, {53, 59}},
	'I':  {{13, 53}, {19, 59}, {33, 39}},
	'J':  {{15, 14, 23, 43, 54, 59, 19}},
	'K':  {{13, 19}, {16, 26, 59}, {26, 53}},
	'L':  {{19, 13, 53}},
	'M':  {{13, 19, 36, 59, 53}},
	'N':  {{13, 19, 53, 59}},
	'O':  {{14, 18, 29, 49, 58, 54, 43, 23, 14}},
	'P':  {{13, 19, 49, 58, 57, 46, 16}},
	'Q':  {{14, 18, 29, 49, 58, 54, 43, 23, 14}, {34, 52}},
	'R':  {{13, 19, 49, 58, 57, 46, 16}, {26, 53}},
	'S':  {{14, 23, 43, 54, 55, 46, 26, 17, 18, 29, 49, 58}},
	'T':  {{19, 59}, {33, 39}},
	'U':  {{19, 14, 23, 43, 54, 59}},
	'V':  {{19, 33, 59}},
	'W':  {{19, 23, 39, 43, 59}},
	'X':  {{13, 59}, {19, 53}},
	'Y':  {{19, 36, 59}, {33, 36}},
	'Z':  {{19, 59, 13, 53}},
	'[':  {{43, 23, 29, 49}},
	'\\': {{19, 53}},
	']':  {{23, 43, 49, 29}},
	'^':  {{17, 39, 57}},
	'_':  {{13, 53}},

	// 0x60
	'`': {{29, 38}},
	'a': {{53, 23, 14, 16, 27, 57, 53}},
	'b': {{17, 47, 56, 54, 43, 13, 19}},
	'c': {{53, 23, 14, 16, 27, 57}},
	'd': {{57, 27, 16, 14, 23, 53, 59}},
	'e': {{15, 55, 56, 47, 27, 16, 14, 23, 53}},
	'f': {{16, 56}, {33, 38, 49, 59}},
	'g': {{21, 41, 52, 57, 27, 16, 15, 24, 54}},
	'h': {{13, 19}, {17, 47, 56, 53}},
	'i': {{13, 53}, {27, 37, 33}, {39}},
	'j': {{11, 31, 42, 47, 27}, {49}},
	'k': {{13, 19}, {15, 35, 57}, {35, 53}},
	'l': {{19, 29, 24, 33, 53}},
	'm': {{13, 17, 47, 56, 53}, {33, 37}},
	'n': {{13, 17, 47, 56, 53}},
	'o': {{14, 16, 27, 47, 56, 54, 43, 23, 14}},
	'p': {{11, 17, 47, 56, 54, 43, 13}},
	'q': {{51, 57, 27, 16, 14, 23, 53}},
	'r': {{13, 17, 47, 56}},
	's': {{13, 43, 54, 45, 25, 16, 27, 57}},
	't': {{17, 57}, {39, 34, 43, 53}},
	'u': {{17, 14, 23, 43, 54, 57}},
	'v': {{17, 33, 57}},
	'w': {{17, 23, 37, 43, 57}},
	'x': {{13, 57}, {17, 53}},
	'y': {{17, 34}, {21, 57}},
	'z': {{17, 57, 13, 53}},
	'{': {{16, 26, 35, 34, 43}, {26, 37, 38, 49}},
	'|': {{33, 39}},
	'}': {{23, 34, 35, 46}, {29, 38, 37, 46, 56}},
	'~': {{16, 27, 45, 56}},
}

var symbolPaths = [256]path.Paths{}

func Centered(str string) path.Paths {
	return alignedText(AlignCenter, path.Point{}, str)
}

func CenteredRow(dx float64, strs ...string) path.Paths {
	return alignedText(AlignCenter, path.Point{X: dx}, strs...)
}

func CenteredColumn(dy float64, strs ...string) path.Paths {
	return alignedText(AlignCenter, path.Point{Y: dy}, strs...)
}

func AlignedColumn(align Align, dy float64, strs ...string) path.Paths {
	return alignedText(align, path.Point{Y: dy}, strs...)
}

func alignedText(align Align, shift path.Point, strs ...string) path.Paths {
	var paths path.Paths

	x0, y0 := -0.5*float64(len(strs)-1)*shift.X, -0.5*float64(len(strs)-1)*shift.Y
	for i, str := range strs {
		i := float64(i)
		p := alignedPaths(align, str).Transform(transform.Move(x0+i*shift.X, y0+i*shift.Y))
		paths = append(paths, p...)
	}

	return paths
}

func alignedPaths(align Align, str string) path.Paths {
	var paths path.Paths

	n := float64(len(str))
	for i, c := range str {
		c := int(c)
		if c >= len(symbolPaths) {
			c = '?'
		}

		t := transform.Move(Width/2+Width*(float64(i)-n*float64(align)), -0.1)
		paths = append(paths, symbolPaths[c].Transform(t)...)
	}

	return paths
}

func init() {
	for i, strokes := range data {
		for _, stroke := range strokes {
			symbolPath := []path.Point{pToXY(stroke[0])}

			for _, point := range stroke[1:] {
				p := pToXY(point)
				symbolPath = append(symbolPath, symbolPath[len(symbolPath)-1], p, p)
			}

			symbolPaths[i] = append(symbolPaths[i], symbolPath)
		}
	}
}

func WeightScale(t transform.T) float64 {
	return min(math.Sqrt(t.Ix*t.Ix+t.Iy*t.Iy), math.Sqrt(t.Jx*t.Jx+t.Jy*t.Jy))
}

func pToXY(p byte) path.Point {
	return path.Point{
		X: (float64(p/10) - 3) * 0.1,
		Y: (float64(p%10) - 5) * 0.1,
	}
}
