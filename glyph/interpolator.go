package glyph

import "temnok/lab/line"

type Interpolator struct {
	*Builder

	started        bool
	startX, startY int
	lastX, lastY   int
}

func (i *Interpolator) AddContourPoint(x, y int) {
	if !i.started {
		i.startX, i.startY = x, y
		i.started = true

		i.Builder.AddContourPoint(x, y)
		return
	}

	if d := i.lastY - y; -1 <= d && d <= 1 {
		i.Builder.AddContourPoint(x, y)
	} else {
		line.Rasterize(i.lastX, i.lastY, x, y, i.Builder.AddContourPoint)
	}

	i.lastX, i.lastY = x, y
}

func (i *Interpolator) FinishContour() {
	if !i.started {
		return
	}

	if i.lastY != i.startY {
		line.Rasterize(i.lastX, i.lastY, i.startX, i.startY, i.Builder.AddContourPoint)
	}
	i.started = false

	i.Builder.FinishContour()
}
