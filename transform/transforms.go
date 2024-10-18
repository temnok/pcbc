package transform

import "math"

var Identity = Transform{Ix: 1, Jy: 1}

func Move(x, y float64) Transform {
	// 1 0
	// 0 1
	// x y
	return Transform{Ix: 1, Jy: 1, Kx: x, Ky: y}
}

func Rotate(ga float64) Transform {
	// cos(a) sin(a)
	// -sin(a) cos(a)
	// 0 0
	ra := ga * math.Pi / 180
	sin, cos := math.Sin(ra), math.Cos(ra)
	
	return Transform{Ix: cos, Iy: sin, Jx: -sin, Jy: cos}
}
