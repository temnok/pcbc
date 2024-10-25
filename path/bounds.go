package path

type Bounds struct {
	initialized bool
	lb, rt      Point
}

func (b *Bounds) addPoint(p Point) {
	if !b.initialized {
		b.initialized = true
		b.lb, b.rt = p, p
		return
	}

	b.lb.X, b.lb.Y = min(b.lb.X, p.X), min(b.lb.Y, p.Y)
	b.rt.X, b.rt.Y = max(b.rt.X, p.X), max(b.rt.Y, p.Y)
}

func (b *Bounds) AddPath(path Path) {
	for _, p := range path {
		b.addPoint(p)
	}
}

func (b *Bounds) AddPaths(paths Paths) {
	for _, p := range paths {
		b.AddPath(p)
	}
}

func (b *Bounds) AddStrokes(strokes Strokes) {
	for d, p := range strokes {
		b1 := &Bounds{}
		b1.AddPaths(p)
		if b1.initialized {
			b.addPoint(Point{X: b1.lb.X - d/2, Y: b1.lb.Y - d/2})
			b.addPoint(Point{X: b1.rt.X + d/2, Y: b1.rt.Y + d/2})
		}
	}
}

func (b *Bounds) Width() float64 {
	return b.rt.X - b.lb.X
}

func (b *Bounds) Height() float64 {
	return b.rt.Y - b.lb.Y
}
