package path

import (
	"math"
	"temnok/pcbc/transform"
)

// Strokes are Paths with added thickness (brush diameter) that serves as a Paths group key.
type Strokes map[float64]Paths

func (strokes Strokes) Append(others ...Strokes) Strokes {
	for _, other := range others {
		for brushD, paths := range other {
			strokes[brushD] = append(strokes[brushD], paths...)
		}
	}

	return strokes
}

func (strokes Strokes) AddPath(brushD float64, path Path) {
	strokes[brushD] = append(strokes[brushD], path)
}

func (strokes Strokes) AddPaths(brushD float64, paths Paths) {
	strokes[brushD] = append(strokes[brushD], paths...)
}

func (strokes Strokes) Apply(t transform.Transform) Strokes {
	res := Strokes{}

	for brushD, paths := range strokes {
		scale := min(math.Sqrt(t.Ix*t.Ix+t.Iy*t.Iy), math.Sqrt(t.Jx*t.Jx+t.Jy*t.Jy))

		newBrushD := scale * brushD
		res[newBrushD] = append(res[newBrushD], paths.Apply(t)...)
	}

	return res
}
