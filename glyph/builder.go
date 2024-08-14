package glyph

type Builder struct {
	addSegment    func(x0, x1, y int)
	started       bool
	x0, x1, y, py int
}

func New(addSegment func(x0, x1, y int)) *Builder {
	return &Builder{
		addSegment: addSegment,
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

	b.push(y)

	b.py = b.y
	b.y = y
	b.x0, b.x1 = x, x
}

func (b *Builder) FinishContour() {
	b.push(b.y)

	b.started = false
}

func (b *Builder) push(y int) {
	b.addSegment(b.x0, b.x1, b.y)
	if y == b.py {
		b.addSegment(b.x0, b.x1, b.y)
	}
}
