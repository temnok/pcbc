package geom

import "math"

const Degree = math.Pi / 180

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

func MoveXY(x, y float64) Transform {
	return Move(XY{x, y})
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

func RotateD(d float64) Transform {
	return Rotate(d * Degree)
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

func ScaleK(k float64) Transform {
	return Scale(XY{k, k})
}
