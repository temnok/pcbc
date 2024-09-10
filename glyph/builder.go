package glyph

import (
	"temnok/lab/line"
)

type Builder struct {
	buf              *buffer
	started          bool
	x0, x1, x, y, py int
	startX, startY   int
}

func (b *Builder) AddContourPoint(x, y int) {
	if !b.started {
		b.x0, b.x1, b.x, b.y, b.py = x, x, x, y, y
		b.startX, b.startY = x, y
		b.started = true
		return
	}

	if d := b.y - y; -1 <= d && d <= 1 {
		b.nextPoint(x, y)
	} else {
		line.Rasterize(b.x, b.y, x, y, b.nextPoint)
	}

	b.x = x
}

func (b *Builder) nextPoint(x, y int) {
	if y == b.y {
		if x < b.x0 {
			b.x0 = x
		} else if b.x1 < x {
			b.x1 = x
		}
		return
	}

	b.addSegments(y)

	b.py = b.y
	b.y = y
	b.x0, b.x1 = x, x
}

func (b *Builder) FinishContour() {
	b.addSegments(b.y)

	if b.x != b.startX || b.y != b.startY {
		line.Rasterize(b.x, b.y, b.x, b.y, b.nextPoint)
	}

	b.started = false
}

func (b *Builder) Rasterize(onRow func(x0, x1, y int)) {
	b.buf.rasterize(onRow)
}

func (b *Builder) Build() *Glyph {
	return b.buf.toGlyph()
}

func (b *Builder) addSegments(y int) {
	if b.buf == nil {
		b.buf = &buffer{}
	}

	b.buf.addSegment(b.x0, b.x1, b.y)

	if y == b.py {
		b.buf.addSegment(b.x0, b.x1, b.y)
	}
}
