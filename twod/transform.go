package twod

import "math"

type Transform struct {
	I, J, K Coord
}

func Identity() Transform {
	// 1 0 0
	// 0 1 0
	return Transform{
		I: Coord{
			X: 1,
			Y: 0,
		},
		J: Coord{
			X: 0,
			Y: 1,
		},
		K: Coord{
			X: 0,
			Y: 0,
		},
	}
}

func Move(p Coord) Transform {
	// 1 0 p.X
	// 0 1 p.Y
	return Transform{
		I: Coord{
			X: 1,
			Y: 0,
		},
		J: Coord{
			X: 0,
			Y: 1,
		},
		K: Coord{
			X: p.X,
			Y: p.Y,
		},
	}
}

func (a Transform) Move(p Coord) Transform {
	return a.Multiply(Move(p))
}

func Rotate(a float64) Transform {
	// cos(a) -sin(a)  0
	// sin(a)  cos(a)  0
	sin := math.Sin(a)
	cos := math.Cos(a)
	return Transform{
		I: Coord{
			X: cos,
			Y: sin,
		},
		J: Coord{
			X: -sin,
			Y: cos,
		},
		K: Coord{
			X: 0,
			Y: 0,
		},
	}
}

func (a Transform) Rotate(b float64) Transform {
	return a.Multiply(Rotate(b))
}

func Scale(p Coord) Transform {
	// p.X   0   0
	//  0   p.Y  0
	return Transform{
		I: Coord{
			X: p.X,
			Y: 0,
		},
		J: Coord{
			X: 0,
			Y: p.Y,
		},
		K: Coord{
			X: 0,
			Y: 0,
		},
	}
}

func (a Transform) Scale(p Coord) Transform {
	return a.Multiply(Scale(p))
}

func ScaleLocked(k float64) Transform {
	return Scale(Coord{k, k})
}

func (a Transform) ScaleLocked(k float64) Transform {
	return a.Multiply(ScaleLocked(k))
}

func (a Transform) Multiply(b Transform) Transform {
	// a.I.X  a.J.X  a.K.X     b.I.X  b.J.X  b.K.X
	// a.I.Y  a.J.Y  a.K.Y  *  b.I.Y  b.J.Y  b.K.Y
	//   0      0      1         0      0      1
	return Transform{
		I: Coord{
			X: a.I.X*b.I.X + a.J.X*b.I.Y, // + a.K.X*0
			Y: a.I.Y*b.I.X + a.J.Y*b.I.Y, // + a.K.Y*0
		},
		J: Coord{
			a.I.X*b.J.X + a.J.X*b.J.Y, // + a.K.X*0
			a.I.Y*b.J.X + a.J.Y*b.J.Y, // + a.K.Y*0
		},
		K: Coord{
			a.I.X*b.K.X + a.J.X*b.K.Y + a.K.X, // *1
			a.I.Y*b.K.X + a.J.Y*b.K.Y + a.K.Y, // *1
		},
	}
}

func (a Transform) Point(b Coord) Coord {
	// a.I.X  a.J.X  a.K.X     b.I.X
	// a.I.Y  a.J.Y  a.K.Y  *  b.I.Y
	//   0      0      1         1
	return Coord{
		a.I.X*b.X + a.J.X*b.Y + a.K.X,
		a.I.Y*b.X + a.J.Y*b.Y + a.K.Y,
	}
}
