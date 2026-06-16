// Copyright © 2025 Alex Temnok. All rights reserved.

package font

import (
	"github.com/temnok/pcbc/path"
	"math"
	"slices"
)

type Align float64

const (
	Width = 0.6 // relative to height 1.0

	XtraLight = 0.0512
	Light     = 0.064
	Medium    = 0.08
	Bold      = 0.1
	XtraBold  = 0.125

	AlignLeft   Align = 0.0
	AlignCenter Align = 0.5
	AlignRight  Align = 1.0
)

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

	for _, paths := range symbolPaths {
		for _, pa := range paths {
			for i, p := range pa {
				pa[i] = path.Point{X: p.X * 0.05, Y: p.Y * 0.05}
			}
		}
	}
}

func pToXY(p byte) path.Point {
	return path.Point{
		X: float64(p%10) - 5,
		Y: 9 - float64(p/10),
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

		if isDiagonal := math.Abs(p1.X-p2.X) == math.Abs(p1.Y-p2.Y); !isDiagonal {
			continue
		}

		sx10, sy10, sx23, sy23 := sign(p1.X-p0.X), sign(p1.Y-p0.Y), sign(p2.X-p3.X), sign(p2.Y-p3.Y)
		sx12, sy12 := sign(p1.X-p2.X), sign(p1.Y-p2.Y)
		if sx12 != 0 && (sx10 == sx12 || sy10 == sy12 || sx23 == -sx12 || sy23 == -sy12) {
			continue
		}

		r := math.Abs(p1.X-p2.X) + 1
		p1.X -= sx10
		p1.Y -= sy10
		p2.X -= sx23
		p2.Y -= sy23

		k := r * 0.55
		if isTerminal := i < 0 || i+9 >= len(p); isTerminal {
			k = r * 0.7
		}

		c1 := path.Point{p1.X + k*sx10, p1.Y + k*sy10}
		c2 := path.Point{p2.X + k*sx23, p2.Y + k*sy23}

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
