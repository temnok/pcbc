// Copyright © 2025 Alex Temnok. All rights reserved.

package transform

type Transform struct {
	Ix, Iy float64
	Jx, Jy float64
	Kx, Ky float64
}

func (t Transform) Multiply(m Transform) Transform {
	return Multiply(t, m)
}

func (t Transform) Move(x, y float64) Transform {
	return t.Multiply(Move(x, y))
}

func (t Transform) Rotate(d float64) Transform {
	return t.Multiply(Rotate(d))
}

func (t Transform) Scale(kx, ky float64) Transform {
	return t.Multiply(Scale(kx, ky))
}

func (t Transform) ScaleK(k float64) Transform {
	return t.Multiply(ScaleK(k))
}

func (t Transform) Apply(x, y float64) (float64, float64) {
	return Apply(x, y, t)
}
