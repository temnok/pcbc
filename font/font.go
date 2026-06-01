// Copyright © 2025 Alex Temnok. All rights reserved.

package font

import (
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
	"math"
	"slices"
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
     11  12  13  14  15  16  17  18  19
     21  22  23  24  25  26  27  28  29
     31  32  33  34  35  36  37  38  39
     41  42  43  44  45  46  47  48  49
     51  52  53  54  55  56  57  58  59
     61  62  63  64  65  66  67  68  69
     71  72  73  74  75  76  77  78  79
     81  82  83  84  85  86  87  88  89
     91  92  93  94  95  96  97  98  99
    101 102 103 104 105 106 107 108 109
    111 112 113 114 115 116 117 118 119
    121 122 123 124 125 126 127 128 129
    131 132 133 134 135 136 137 138 139
    141 142 143 144 145 146 147 148 149
    151 152 153 154 155 156 157 158 159
    161 162 163 164 165 166 167 168 169
    171 172 173 174 175 176 177 178 179
*/

var data = [][][]byte{
	// 0x20
	' ':  {},
	'!':  {{15, 95}, {135}},
	'"':  {{13, 53}, {17, 57}},
	'#':  {{13, 133}, {17, 137}, {51, 59}, {91, 99}},
	'$':  {{15, 175}, {59, 37, 33, 51, 71, 93, 97, 119, 139, 157, 153, 131}},
	'%':  {{11, 13, 33, 31, 11}, {19, 131}, {117, 119, 139, 137, 117}},
	'&':  {{37, 15, 13, 31, 139}, {73, 91, 111, 133, 135, 99}},
	'\'': {{15, 55}},
	'(':  {{15, 53, 93, 135}},
	')':  {{15, 57, 97, 135}},
	'*':  {{35, 115}, {51, 99}, {59, 91}},
	'+':  {{35, 115}, {71, 79}},
	',':  {{135, 173}},
	'-':  {{71, 79}},
	'.':  {{135}},
	'/':  {{19, 131}},

	// 0x30
	'0': {{15, 17, 39, 119, 137, 133, 111, 31, 13, 15}, {59, 57, 93, 91}},
	'1': {{51, 15, 135}, {131, 139}},
	'2': {{31, 13, 17, 39, 59, 131, 139}},
	'3': {{31, 13, 17, 39, 59, 77, 75, 77, 99, 119, 137, 133, 111}},
	'4': {{15, 91, 99}, {59, 139}},
	'5': {{19, 11, 71, 77, 99, 119, 137, 131}},
	'6': {{15, 71, 111, 133, 137, 119, 99, 77, 71}},
	'7': {{11, 19, 133}},
	'8': {{75, 73, 51, 31, 13, 17, 39, 59, 77, 73, 91, 111, 133, 137, 119, 99, 77, 75}},
	'9': {{79, 73, 51, 31, 13, 17, 39, 79, 135}},
	':': {{35}, {135}},
	';': {{35}, {135, 173}},
	'<': {{19, 71, 139}},
	'=': {{51, 59}, {91, 99}},
	'>': {{11, 79, 131}},
	'?': {{31, 13, 17, 39, 59, 95}, {135}},

	// 0x40
	'@': {{59, 55, 95, 99, 39, 17, 13, 31, 111, 133, 137}},
	'A': {{71, 79}, {131, 31, 13, 17, 39, 139}},
	'B': {{71, 77, 59, 39, 17, 11, 131, 137, 119, 99, 77, 75}},
	'C': {{39, 17, 13, 31, 111, 133, 137, 119}},
	'D': {{11, 17, 39, 119, 137, 131, 11}},
	'E': {{19, 11, 131, 139}, {71, 77}},
	'F': {{19, 11, 131}, {71, 77}},
	'G': {{39, 17, 15, 13, 31, 111, 133, 137, 119, 79, 75}},
	'H': {{11, 131}, {19, 139}, {71, 79}},
	'I': {{11, 19}, {15, 135}, {131, 139}},
	'J': {{11, 19, 119, 137, 133, 111, 91}},
	'K': {{11, 131}, {19, 73, 71}, {73, 139}},
	'L': {{11, 131, 139}},
	'M': {{131, 11, 75, 19, 139}},
	'N': {{19, 139, 11, 131}},
	'O': {{15, 17, 39, 119, 137, 133, 111, 31, 13, 15}},

	// 0x50
	'P':  {{71, 77, 59, 39, 17, 11, 131}},
	'Q':  {{15, 17, 39, 119, 137, 133, 111, 31, 13, 15}, {115, 159}},
	'R':  {{71, 77, 59, 39, 17, 11, 131}, {73, 139}},
	'S':  {{39, 17, 13, 31, 51, 73, 77, 99, 119, 137, 133, 111}},
	'T':  {{11, 19}, {15, 135}},
	'U':  {{11, 111, 133, 137, 119, 19}},
	'V':  {{11, 135, 19}},
	'W':  {{11, 133, 15, 137, 19}},
	'X':  {{11, 139}, {19, 131}},
	'Y':  {{11, 75, 19}, {75, 135}},
	'Z':  {{11, 19, 131, 139}},
	'[':  {{17, 13, 133, 137}},
	'\\': {{11, 139}},
	']':  {{13, 17, 137, 133}},
	'^':  {{51, 15, 59}},
	'_':  {{131, 139}},

	// 0x60
	'`': {{13, 35}},
	'a': {{59, 139, 133, 111, 71, 53, 59}},
	'b': {{11, 131, 137, 119, 79, 57, 51}},
	'c': {{59, 53, 71, 111, 133, 139}},
	'd': {{19, 139, 133, 111, 71, 53, 59}},
	'e': {{91, 99, 79, 57, 53, 71, 111, 133, 139}},
	'f': {{19, 17, 35, 135}, {71, 79}},
	'g': {{119, 113, 91, 71, 53, 59, 159, 177, 173}},
	'h': {{11, 131}, {51, 57, 79, 139}},
	'i': {{15}, {53, 55, 135}, {131, 139}},
	'j': {{17}, {53, 57, 157, 175, 171}},
	'k': {{11, 131}, {59, 95, 91}, {95, 139}},
	'l': {{11, 13, 113, 135, 139}},
	'm': {{55, 135}, {131, 51, 57, 79, 139}},
	'n': {{131, 51, 57, 79, 139}},
	'o': {{55, 57, 79, 119, 137, 133, 111, 71, 53, 55}},

	// 0x70
	'p': {{131, 137, 119, 79, 57, 51, 171}},
	'q': {{139, 133, 111, 71, 53, 59, 179}},
	'r': {{79, 57, 51, 131}},
	's': {{59, 53, 71, 93, 97, 119, 137, 131}},
	't': {{15, 115, 137, 139}, {51, 59}},
	'u': {{51, 111, 133, 137, 119, 59}},
	'v': {{51, 135, 59}},
	'w': {{51, 133, 55, 137, 59}},
	'x': {{51, 139}, {59, 131}},
	'y': {{51, 115}, {59, 173}},
	'z': {{51, 59, 131, 139}},
	'{': {{17, 35, 55, 73, 71, 73, 95, 115, 137}},
	'|': {{15, 135}},
	'}': {{13, 35, 55, 77, 79, 77, 95, 115, 133}},
	'~': {{71, 53, 97, 79}},
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

			roundedPath := slices.Clone(symbolPath)
			//roundPath(roundedPath, symbolPath)

			symbolPaths[i] = append(symbolPaths[i], roundedPath)
		}
	}
}

var circleK = 4 * (math.Sqrt(2) - 1) / 3

func roundPath(dst, p []path.Point) {
	for i := -3; i+6 < len(p); i += 3 {
		p1, c1, c2, p2 := p[i+3], p[i+4], p[i+5], p[i+6]

		p0 := p1
		if i >= 0 {
			p0 = p[i]
		}

		p3 := p2
		if i+9 < len(p) {
			p3 = p[i+9]
		}

		if !equal(abs(p1.X-p2.X), 0.1) || !equal(abs(p1.Y-p2.Y), 0.1) ||
			(p0.X != p1.X && p0.Y != p1.Y) || (p2.X != p3.X && p2.Y != p3.Y) {
			continue
		}

		var sx10, sy10, sx23, sy23 = sign(p1.X - p0.X), sign(p1.Y - p0.Y), sign(p2.X - p3.X), sign(p2.Y - p3.Y)

		const k = 0.05
		p1.X -= k * sx10
		p1.Y -= k * sy10
		p2.X -= k * sx23
		p2.Y -= k * sy23

		var c = circleK * 0.75 * 0.1
		c1.X += c * sx10
		c1.Y += c * sy10
		c2.X += c * sx23
		c2.Y += c * sy23

		//dst[i+2] = p1
		dst[i+3] = p1
		dst[i+4] = c1
		dst[i+5] = c2
		dst[i+6] = p2
		//dst[i+7] = p2
	}
}

func pToXY(p byte) path.Point {
	return path.Point{
		X: (float64(p%10) - 5) * 0.05,
		Y: (9 - float64(p/10)) * 0.05,
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}

	return a
}

func equal(a, b float64) bool {
	return abs(a-b) < 1e-9
}

func sign(a float64) float64 {
	switch {
	case a < 0:
		return -1
	case a == 0:
		return 0
	default:
		return 1
	}
}
