package font

import (
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

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
