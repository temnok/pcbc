// Copyright © 2025 Alex Temnok. All rights reserved.

package font

import (
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
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
	'"':  {{13, 43}, {17, 47}},
	'#':  {{13, 133}, {17, 137}, {51, 59}, {91, 99}},
	'$':  {{15, 175}, {59, 37, 33, 51, 71, 93, 97, 119, 139, 157, 153, 131}},
	'%':  {{13, 14, 25, 45, 54, 52, 41, 21, 12, 13}, {19, 131}, {97, 98, 109, 129, 138, 136, 125, 105, 96, 97}},
	'&':  {{35, 24, 22, 31, 51, 73, 129}, {73, 82, 91, 111, 133, 135, 126, 89}},
	'\'': {{15, 45}},
	'(':  {{17, 53, 93, 137}},
	')':  {{13, 57, 97, 133}},
	'*':  {{35, 115}, {51, 99}, {59, 91}},
	'+':  {{35, 115}, {71, 79}},
	',':  {{135, 164}},
	'-':  {{71, 79}},
	'.':  {{135}},
	'/':  {{19, 131}},

	// 0x30
	'0': {{15, 17, 39, 119, 137, 133, 111, 31, 13, 15}, {59, 57, 93, 91}},
	'1': {{51, 15, 135}, {131, 139}},
	'2': {{31, 13, 17, 39, 59, 131, 139}},
	'3': {{31, 13, 17, 39, 59, 77, 75, 77, 99, 119, 137, 133, 111}},
	'4': {{16, 91, 99}, {59, 139}},
	'5': {{19, 11, 71, 77, 99, 119, 137, 131}},
	'6': {{18, 16, 61, 71, 111, 133, 137, 119, 99, 77, 71}},
	'7': {{11, 19, 133}},
	'8': {{75, 73, 51, 31, 13, 17, 39, 59, 77, 73, 91, 111, 133, 137, 119, 99, 77, 75}},
	'9': {{79, 73, 51, 31, 13, 17, 39, 89, 134, 132}},
	':': {{35}, {135}},
	';': {{35}, {135, 164}},
	'<': {{19, 71, 139}},
	'=': {{51, 59}, {91, 99}},
	'>': {{11, 79, 131}},
	'?': {{31, 13, 17, 39, 49, 85, 95}, {135}},

	// 0x40
	'@': {{59, 55, 55, 95, 95, 99, 99, 39, 17, 13, 31, 111, 133, 137}},
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
	'K': {{11, 131}, {73, 71}, {19, 73, 139}},
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
	'a': {{61, 52, 58, 69, 139}, {89, 98, 92, 101, 121, 132, 138, 129}},
	'b': {{11, 131}, {61, 52, 57, 79, 119, 137, 132, 121}},
	'c': {{69, 58, 53, 71, 111, 133, 138, 129}},
	'd': {{19, 139}, {69, 58, 53, 71, 111, 133, 138, 129}},
	'e': {{91, 99, 79, 57, 53, 71, 111, 133, 138, 129}},
	'f': {{19, 17, 35, 135}, {61, 69}},
	'g': {{59, 159, 177, 173, 162}, {69, 58, 53, 71, 101, 123, 128, 119}},
	'h': {{11, 131}, {61, 52, 57, 79, 139}},
	'i': {{15}, {52, 55, 135}, {132, 138}},
	'j': {{16}, {52, 56, 156, 174, 171}},
	'k': {{11, 131}, {58, 94, 138}, {91, 94}},
	'l': {{11, 14, 114, 136, 139}},
	'm': {{51, 131}, {61, 52, 54, 65, 135}, {65, 56, 58, 69, 139}},
	'n': {{51, 131}, {61, 52, 57, 79, 139}},
	'o': {{55, 57, 79, 119, 137, 133, 111, 71, 53, 55}},

	// 0x70
	'p': {{51, 171}, {61, 52, 57, 79, 119, 137, 132, 121}},
	'q': {{59, 179}, {69, 58, 53, 71, 111, 133, 138, 129}},
	'r': {{51, 131}, {61, 52, 57, 79}},
	's': {{69, 58, 52, 61, 81, 92, 98, 109, 129, 138, 132, 121}},
	't': {{15, 115, 137, 139}, {51, 59}},
	'u': {{51, 111, 133, 137, 119, 59}},
	'v': {{51, 135, 59}},
	'w': {{51, 133, 55, 137, 59}},
	'x': {{51, 139}, {59, 131}},
	'y': {{51, 135}, {59, 173}},
	'z': {{51, 59, 131, 139}},
	'{': {{17, 35, 55, 73, 72, 73, 95, 115, 137}},
	'|': {{15, 135}},
	'}': {{13, 35, 55, 77, 78, 77, 95, 115, 133}},
	'~': {{71, 62, 63, 64, 86, 87, 88, 79}},
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
			roundPath(roundedPath, symbolPath)

			symbolPaths[i] = append(symbolPaths[i], roundedPath)
		}
	}
}

func pToXY(p byte) path.Point {
	return path.Point{
		X: (float64(p%10) - 5) * 0.05,
		Y: (9 - float64(p/10)) * 0.05,
	}
}

func roundPath(dst, p []path.Point) {
	for i := -3; i+6 < len(p); i += 3 {
		p1, p2 := p[i+3], p[i+6]

		p0, p3 := p1, p2
		if i >= 0 {
			p0 = p[i]
		}
		if i+9 < len(p) {
			p3 = p[i+9]
		}

		if !equal(abs(p1.X-p2.X), abs(p1.Y-p2.Y)) {
			continue
		}

		sx10, sy10, sx23, sy23 := sign(p1.X-p0.X), sign(p1.Y-p0.Y), sign(p2.X-p3.X), sign(p2.Y-p3.Y)
		sx12, sy12 := sign(p1.X-p2.X), sign(p1.Y-p2.Y)
		if sx12 != 0 && (sx10 == sx12 || sy10 == sy12 || sx23 == -sx12 || sy23 == -sy12) {
			continue
		}

		r := abs(p1.X-p2.X) + 0.05

		k := 0.05
		p1.X -= k * sx10
		p1.Y -= k * sy10
		p2.X -= k * sx23
		p2.Y -= k * sy23

		c := r * 0.55
		if isTerminal := i < 0 || i+9 >= len(p); isTerminal {
			c = r * 0.75
		}

		c1 := path.Point{p1.X + c*sx10, p1.Y + c*sy10}
		c2 := path.Point{p2.X + c*sx23, p2.Y + c*sy23}

		if i+2 >= 0 {
			dst[i+2] = p1
		}

		dst[i+3] = p1
		dst[i+4] = c1
		dst[i+5] = c2
		dst[i+6] = p2

		if i+7 < len(dst) {
			dst[i+7] = p2
		}
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
