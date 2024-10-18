package transform

type Transform struct {
	Ix, Iy float64
	Jx, Jy float64
	Kx, Ky float64
}

func (a Transform) Multiply(b Transform) Transform {
	// a.Ix  a.Iy  0     b.Ix  b.Iy  0
	// a.Jx  a.Jy  0  *  b.Jx  b.Jy  0
	// a.Kx  a.Ky  1     b.Kx  b.Ky  1

	return Transform{
		Ix: a.Ix*b.Ix + a.Iy*b.Jx, Iy: a.Ix*b.Iy + a.Iy*b.Jy,
		Jx: a.Jx*b.Ix + a.Jy*b.Jx, Jy: a.Jx*b.Iy + a.Jy*b.Jy,
		Kx: a.Kx*b.Ix + a.Ky*b.Jx + b.Kx, Ky: a.Kx*b.Iy + a.Ky*b.Jy + b.Ky,
	}
}
