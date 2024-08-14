package glyph

type Builder struct {
	buf           *buffer
	started       bool
	x0, x1, y, py int
}

func New() *Builder {
	return &Builder{
		buf: &buffer{},
	}
}

func (b *Builder) AddPoint(x, y int) {
	if !b.started {
		b.x0, b.x1, b.y, b.py = x, x, y, y
		b.started = true
		return
	}

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

	b.started = false
}

func (b *Builder) Build() *Glyph {
	return b.buf.toGlyph()
}

func (b *Builder) addSegments(y int) {
	b.buf.addSegment(b.x0, b.x1, b.y)
	if y == b.py {
		b.buf.addSegment(b.x0, b.x1, b.y)
	}
}
