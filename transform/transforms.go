package transform

import "math"

var Identity = Transform{Ix: 1, Jy: 1}

func Apply(x, y float64, t Transform) (float64, float64) {
	//            Ix Iy 0
	//  x y 1  *  Jx Jy 0
	//            Kx Ky 1
	return x*t.Ix + y*t.Jx + t.Kx,
		x*t.Iy + y*t.Jy + t.Ky
}

func Multiply(a, b Transform) Transform {
	// a.Ix  a.Iy  0     b.Ix  b.Iy  0
	// a.Jx  a.Jy  0  *  b.Jx  b.Jy  0
	// a.Kx  a.Ky  1     b.Kx  b.Ky  1
	return Transform{
		Ix: a.Ix*b.Ix + a.Iy*b.Jx, Iy: a.Ix*b.Iy + a.Iy*b.Jy,
		Jx: a.Jx*b.Ix + a.Jy*b.Jx, Jy: a.Jx*b.Iy + a.Jy*b.Jy,
		Kx: a.Kx*b.Ix + a.Ky*b.Jx + b.Kx, Ky: a.Kx*b.Iy + a.Ky*b.Jy + b.Ky,
	}
}

func Move(x, y float64) Transform {
	// 1 0 0
	// 0 1 0
	// x y 1
	return Transform{Ix: 1, Jy: 1, Kx: x, Ky: y}
}

func Rotate(d float64) Transform {
	//  cos(a) sin(a) 0
	// -sin(a) cos(a) 0
	//    0     0     1
	r := d * math.Pi / 180
	sin, cos := math.Sin(r), math.Cos(r)

	return Transform{Ix: cos, Iy: sin, Jx: -sin, Jy: cos}
}

func Scale(x, y float64) Transform {
	// x 0 0
	// 0 y 0
	// 0 0 1
	return Transform{Ix: x, Jy: y}
}

func ScaleK(k float64) Transform {
	return Scale(k, k)
}
