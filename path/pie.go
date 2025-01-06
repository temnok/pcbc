// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"math"
	"temnok/pcbc/transform"
)

func Pie(n int, r1, r2, a1 float64) Paths {
	var res Paths

	angle := 360 / float64(n)
	for i := 0; i < n; i++ {
		a := a1/2 + float64(i)*angle
		c := PiePiece(r1, r2, angle-a1).Apply(transform.Rotate(a))
		res = append(res, c)
	}

	return res
}

func PiePiece(r1, r2, ga float64) Path {
	ra := math.Pi * ga / 180

	a1 := Point{r1, 0}
	b1 := Point{r1 * math.Cos(ra), r1 * math.Sin(ra)}
	p11, p12 := arc(a1, b1)

	a2 := Point{r2, 0}
	b2 := Point{r2 * math.Cos(ra), r2 * math.Sin(ra)}
	p21, p22 := arc(a2, b2)

	return Path{
		a1, a1,
		a2, a2, p21,
		p22, b2, b2,
		b1, b1, p12,
		p11, a1,
	}
}

func arc(a, b Point) (p1, p2 Point) {
	q1 := a.X*a.X + a.Y*a.Y
	q2 := q1 + a.X*b.X + a.Y*b.Y

	//k2 := (4.0 / 3.0) * (math.Sqrt(2*q1*q2) - q2) / (a.X*b.Y - a.Y*b.X)
	k2 := (4.0 / 3.0) * (a.X*b.Y - a.Y*b.X) / (math.Sqrt(2*q1*q2) + q2)

	return Point{a.X - k2*a.Y, a.Y + k2*a.X}, Point{b.X + k2*b.Y, b.Y - k2*b.X}
}
