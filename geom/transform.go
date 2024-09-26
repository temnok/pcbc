package geom

type Transform struct {
	I, J, K XY
}

func (a Transform) Move(p XY) Transform {
	return a.Multiply(Move(p))
}

func (a Transform) MoveXY(x, y float64) Transform {
	return a.Multiply(MoveXY(x, y))
}

func (a Transform) Rotate(b float64) Transform {
	return a.Multiply(Rotate(b))
}

func (a Transform) RotateD(d float64) Transform {
	return a.Multiply(RotateD(d))
}

func (a Transform) Scale(p XY) Transform {
	return a.Multiply(Scale(p))
}

func (a Transform) ScaleK(k float64) Transform {
	return a.Multiply(ScaleK(k))
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

func (a Transform) PointXY(x, y float64) XY {
	return a.Point(XY{x, y})
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
