package geom

import "math"

type Transform struct {
	I, J, K XY
}

func Identity() Transform {
	// 1 0 0
	// 0 1 0
	return Transform{
		I: XY{
			X: 1,
			Y: 0,
		},
		J: XY{
			X: 0,
			Y: 1,
		},
		K: XY{
			X: 0,
			Y: 0,
		},
	}
}

func Move(p XY) Transform {
	// 1 0 p.X
	// 0 1 p.Y
	return Transform{
		I: XY{
			X: 1,
			Y: 0,
		},
		J: XY{
			X: 0,
			Y: 1,
		},
		K: XY{
			X: p.X,
			Y: p.Y,
		},
	}
}

func (a Transform) Move(p XY) Transform {
	return a.Multiply(Move(p))
}

func Rotate(a float64) Transform {
	// cos(a) -sin(a)  0
	// sin(a)  cos(a)  0
	sin := math.Sin(a)
	cos := math.Cos(a)
	return Transform{
		I: XY{
			X: cos,
			Y: sin,
		},
		J: XY{
			X: -sin,
			Y: cos,
		},
		K: XY{
			X: 0,
			Y: 0,
		},
	}
}

func (a Transform) Rotate(b float64) Transform {
	return a.Multiply(Rotate(b))
}

func Scale(p XY) Transform {
	// p.X   0   0
	//  0   p.Y  0
	return Transform{
		I: XY{
			X: p.X,
			Y: 0,
		},
		J: XY{
			X: 0,
			Y: p.Y,
		},
		K: XY{
			X: 0,
			Y: 0,
		},
	}
}

func (a Transform) Scale(p XY) Transform {
	return a.Multiply(Scale(p))
}

func ScaleLocked(k float64) Transform {
	return Scale(XY{k, k})
}

func (a Transform) ScaleLocked(k float64) Transform {
	return a.Multiply(ScaleLocked(k))
}

func (a Transform) Multiply(b Transform) Transform {
	// a.I.X  a.J.X  a.K.X     b.I.X  b.J.X  b.K.X
	// a.I.Y  a.J.Y  a.K.Y  *  b.I.Y  b.J.Y  b.K.Y
	//   0      0      1         0      0      1
	return Transform{
		I: XY{
			X: a.I.X*b.I.X + a.J.X*b.I.Y, // + a.K.X*0
			Y: a.I.Y*b.I.X + a.J.Y*b.I.Y, // + a.K.Y*0
		},
		J: XY{
			a.I.X*b.J.X + a.J.X*b.J.Y, // + a.K.X*0
			a.I.Y*b.J.X + a.J.Y*b.J.Y, // + a.K.Y*0
		},
		K: XY{
			a.I.X*b.K.X + a.J.X*b.K.Y + a.K.X, // *1
			a.I.Y*b.K.X + a.J.Y*b.K.Y + a.K.Y, // *1
		},
	}
}

func (a Transform) Point(b XY) XY {
	// a.I.X  a.J.X  a.K.X     b.I.X
	// a.I.Y  a.J.Y  a.K.Y  *  b.I.Y
	//   0      0      1         1
	return XY{
		a.I.X*b.X + a.J.X*b.Y + a.K.X,
		a.I.Y*b.X + a.J.Y*b.Y + a.K.Y,
	}
}
