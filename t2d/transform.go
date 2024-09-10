package t2d

import "math"

type Vector [2]float64

type Transform [3]Vector

func Identity() Transform {
	// 1 0 0
	// 0 1 0
	return Transform{
		{
			1,
			0,
		},
		{
			0,
			1,
		},
		{
			0,
			0,
		},
	}
}

func Move(v Vector) Transform {
	// 1 0 v[0]
	// 0 1 v[1]
	return Transform{
		{
			1,
			0,
		},
		{
			0,
			1,
		},
		{
			v[0],
			v[1],
		},
	}
}

func (a Transform) Move(b Vector) Transform {
	return a.Multiply(Move(b))
}

func Rotate(a float64) Transform {
	// cos(a) -sin(a)  0
	// sin(a)  cos(a)  0
	sin := math.Sin(a)
	cos := math.Cos(a)
	return Transform{
		{
			cos,
			sin,
		},
		{
			-sin,
			cos,
		},
		{
			0,
			0,
		},
	}
}

func (a Transform) Rotate(b float64) Transform {
	return a.Multiply(Rotate(b))
}

func Scale(v Vector) Transform {
	// v[0]  0    0
	//  0   v[1]  0
	return Transform{
		{
			v[0],
			0,
		},
		{
			0,
			v[1],
		},
		{
			0,
			0,
		},
	}
}

func (a Transform) Scale(b Vector) Transform {
	return a.Multiply(Scale(b))
}

func ScaleLocked(k float64) Transform {
	return Scale(Vector{k, k})
}

func (a Transform) ScaleLocked(k float64) Transform {
	return a.Multiply(ScaleLocked(k))
}

func (a Transform) Multiply(b Transform) Transform {
	// a[0][0] a[1][0] a[2][0]     b[0][0] b[1][0] b[2][0]
	// a[0][1] a[1][1] a[2][1]  *  b[0][1] b[1][1] b[2][1]
	//    0       0       1           0       0       1
	return Transform{
		{
			a[0][0]*b[0][0] + a[1][0]*b[0][1], // + a[2][0]*0
			a[0][1]*b[0][0] + a[1][1]*b[0][1], // + a[2][1]*0
		},
		{
			a[0][0]*b[1][0] + a[1][0]*b[1][1], // + a[2][0]*0
			a[0][1]*b[1][0] + a[1][1]*b[1][1], // + a[2][1]*0
		},
		{
			a[0][0]*b[2][0] + a[1][0]*b[2][1] + a[2][0], // *1
			a[0][1]*b[2][0] + a[1][1]*b[2][1] + a[2][1], // *1
		},
	}
}

func (a Transform) Point(b Vector) Vector {
	// a[0][0] a[1][0] a[2][0]     b[0][0]
	// a[0][1] a[1][1] a[2][1]  *  b[0][1]
	//    0       0       1           1
	return Vector{
		a[0][0]*b[0] + a[1][0]*b[1] + a[2][0],
		a[0][1]*b[0] + a[1][1]*b[1] + a[2][1],
	}
}
