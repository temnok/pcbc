package font

import (
	"math"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type Align float64

const (
	Width = 0.65 // relative to height 1.0

	weightScale = 1.0 / 5000.0
	Normal      = 400 * weightScale
	SemiBold    = 600 * weightScale
	Bold        = 700 * weightScale

	AlignLeft   Align = 0.0
	AlignCenter Align = 0.5
	AlignRight  Align = 1.0
)

/*

Vector (linear) font matrix. Data below uses its point numbers.

11  12  13  14  15

21  22  23  24  25

31  32  33  34  35

41  42  43  44  45

51  52  53  54  55

61  62  63  64  65

71  72  73  74  75

81  82  83  84  85

91  92  93  94  95

*/

var data = [][][]byte{
	'!':  {{13, 53}, {73}},
	'"':  {{12, 32}, {14, 34}},
	'#':  {{12, 72}, {14, 74}, {31, 35}, {51, 55}},
	'$':  {{71, 82, 84, 75, 65, 54, 52, 41, 31, 22, 24, 35}, {13, 93}},
	'%':  {{11, 12, 22, 21, 11}, {15, 71}, {64, 65, 75, 74, 64}},
	'&':  {{75, 21, 12, 13, 24, 51, 61, 72, 73, 55}},
	'\'': {{13, 33}},
	'(':  {{73, 52, 32, 13}},
	')':  {{13, 34, 54, 73}},
	'*':  {{13, 53}, {21, 45}, {41, 25}},
	'+':  {{23, 63}, {41, 45}},
	',':  {{73, 92}},
	'-':  {{41, 45}},
	'.':  {{73}},
	'/':  {{15, 71}},

	'0': {{61, 21, 12, 14, 25, 65, 74, 72, 61}, {51, 52, 34, 35}},
	'1': {{31, 13, 73}, {71, 75}},
	'2': {{21, 12, 14, 25, 35, 71, 75}},
	'3': {{21, 12, 14, 25, 35, 44, 43, 44, 55, 65, 74, 72, 61}},
	'4': {{13, 51, 55}, {35, 75}},
	'5': {{15, 11, 41, 44, 55, 65, 74, 71}},
	'6': {{13, 41, 61, 72, 74, 65, 55, 44, 41}},
	'7': {{11, 15, 72}},
	'8': {{42, 31, 21, 12, 14, 25, 35, 44, 42, 51, 61, 72, 74, 65, 55, 44}},
	'9': {{45, 42, 31, 21, 12, 14, 25, 45, 73}},

	':': {{23}, {73}},
	';': {{23}, {73, 92}},
	'<': {{75, 41, 15}},
	'=': {{31, 35}, {51, 55}},
	'>': {{11, 45, 71}},
	'?': {{21, 12, 14, 25, 35, 53}, {73}},

	'@': {{74, 72, 61, 21, 12, 14, 25, 55, 53, 33, 35}},
	'A': {{71, 21, 12, 14, 25, 75}, {41, 45}},
	'B': {{41, 44, 35, 25, 14, 11, 71, 74, 65, 55, 44}},
	'C': {{25, 14, 12, 21, 61, 72, 74, 65}},
	'D': {{71, 11, 14, 25, 65, 74, 71}},
	'E': {{75, 71, 11, 15}, {41, 44}},
	'F': {{71, 11, 15}, {41, 44}},
	'G': {{43, 45, 65, 74, 72, 61, 21, 12, 13, 14, 25}},
	'H': {{11, 71}, {15, 75}, {41, 45}},
	'I': {{13, 73}, {11, 15}, {71, 75}},
	'J': {{11, 15, 65, 74, 72, 61, 51}},
	'K': {{11, 71}, {41, 42, 15}, {42, 75}},
	'L': {{11, 71, 75}},
	'M': {{71, 11, 43, 15, 75}},
	'N': {{71, 11, 75, 15}},
	'O': {{61, 21, 12, 14, 25, 65, 74, 72, 61}},
	'P': {{71, 11, 14, 25, 35, 44, 41}},
	'Q': {{61, 21, 12, 14, 25, 65, 74, 72, 61}, {63, 85}},
	'R': {{71, 11, 14, 25, 35, 44, 41}, {42, 75}},
	'S': {{25, 14, 12, 21, 31, 42, 44, 55, 65, 74, 72, 61}},
	'T': {{11, 15}, {13, 73}},
	'U': {{11, 61, 72, 74, 65, 15}},
	'V': {{11, 73, 15}},
	'W': {{11, 72, 13, 74, 15}},
	'X': {{11, 75}, {15, 71}},
	'Y': {{11, 43, 15}, {43, 73}},
	'Z': {{11, 15, 71, 75}},

	'[':  {{74, 72, 12, 14}},
	'\\': {{11, 75}},
	']':  {{12, 14, 74, 72}},
	'^':  {{31, 13, 35}},
	'_':  {{71, 75}},

	'`': {{12, 23}},
	'a': {{65, 74, 72, 61, 41, 32, 35, 75}},
	'b': {{31, 34, 45, 65, 74, 71, 11}},
	'c': {{75, 72, 61, 41, 32, 35}},
	'd': {{15, 75, 72, 61, 41, 32, 35}},
	'e': {{75, 72, 61, 41, 32, 34, 45, 55, 51}},
	'f': {{15, 14, 23, 73}, {41, 45}},
	'g': {{75, 72, 61, 41, 32, 35, 85, 94, 92}},
	'h': {{11, 71}, {31, 34, 45, 75}},
	'i': {{32, 33, 73}, {71, 75}, {13}},
	'j': {{32, 34, 84, 93, 91}, {14}},
	'k': {{11, 71}, {51, 53, 35}, {53, 75}},
	'l': {{11, 12, 62, 73, 75}},
	'm': {{71, 31, 34, 45, 75}, {33, 73}},
	'n': {{71, 31, 34, 45, 75}},
	'o': {{61, 41, 32, 34, 45, 65, 74, 72, 61}},
	'p': {{91, 31, 34, 45, 65, 74, 71}},
	'q': {{75, 72, 61, 41, 32, 35, 95}},
	'r': {{71, 31}, {41, 32, 34, 45}},
	's': {{71, 74, 65, 54, 52, 41, 32, 35}},
	't': {{13, 63, 74, 75}, {31, 35}},
	'u': {{65, 74, 72, 61, 31}, {35, 75}},
	'v': {{31, 73, 35}},
	'w': {{31, 72, 33, 74, 35}},
	'x': {{31, 75}, {35, 71}},
	//'y': {{31, 63}, {35, 91}},
	'y': {{35, 85, 94, 92}, {75, 72, 61, 31}},
	'z': {{31, 35, 71, 75}},

	'{': {{74, 63, 53, 42, 41}, {42, 33, 23, 14}},
	'|': {{13, 73}},
	'}': {{12, 23, 33, 44, 45}, {44, 53, 63, 72}},
	'~': {{41, 32, 54, 45}},
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
		p := alignedPaths(align, str).Apply(transform.Move(x0+i*shift.X, y0+i*shift.Y))
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

		t := transform.Move(Width*(float64(i)-n*float64(align)), 0.4)
		paths = append(paths, symbolPaths[c].Apply(t)...)
	}

	return paths
}

func init() {
	for i, paths := range data {
		symbolPaths[i] = make(path.Paths, len(paths))

		for j, p := range paths {
			a := pToXY(p[0])
			symbolPaths[i][j] = []path.Point{a}

			for _, p := range p[1:] {
				b := pToXY(p)
				symbolPaths[i][j] = append(symbolPaths[i][j], a, b, b)

				a = b
			}
		}
	}
}

func WeightScale(t transform.Transform) float64 {
	return min(math.Sqrt(t.Ix*t.Ix+t.Iy*t.Iy), math.Sqrt(t.Jx*t.Jx+t.Jy*t.Jy))
}

func pToXY(p byte) path.Point {
	return path.Point{
		X: float64(p%10) / 10.0,
		Y: -float64(p/10) / 10.0,
	}
}
