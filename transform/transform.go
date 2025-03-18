// Copyright © 2025 Alex Temnok. All rights reserved.

package transform

type T struct {
	Ix, Iy float64
	Jx, Jy float64
	Kx, Ky float64
}

var I = T{Ix: 1, Jy: 1}

func (t T) Det() float64 {
	return t.Ix*t.Jy - t.Iy*t.Jx
}

func (t T) Multiply(m T) T {
	return Multiply(t, m)
}

func (t T) Move(x, y float64) T {
	return Multiply(t, Move(x, y))
}

func (t T) RotateDegrees(d float64) T {
	return Multiply(t, RotateDegrees(d))
}

func (t T) Scale(kx, ky float64) T {
	return Multiply(t, Scale(kx, ky))
}

func (t T) ScaleUniformly(k float64) T {
	return Multiply(t, ScaleUniformly(k))
}

func (t T) Apply(x, y float64) (float64, float64) {
	return Apply(x, y, t)
}
