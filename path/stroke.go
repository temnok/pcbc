package path

import (
	"math"
	"temnok/lab/geom"
)

// Strokes are Paths with added thickness (brush diameter) that serves as a Paths group key.
type Strokes map[float64]Paths

func (strokes Strokes) Merge(other Strokes) Strokes {
	res := Strokes{}

	res.Add(strokes)
	res.Add(other)

	return res
}

func (strokes Strokes) Add(s Strokes) {
	for brushD, paths := range s {
		strokes[brushD] = append(strokes[brushD], paths...)
	}
}

func (strokes Strokes) AddPath(brushD float64, path Path) {
	strokes[brushD] = append(strokes[brushD], path)
}

func (strokes Strokes) AddPaths(brushD float64, paths Paths) {
	strokes[brushD] = append(strokes[brushD], paths...)
}

func (strokes Strokes) Transform(t geom.Transform) Strokes {
	res := Strokes{}

	for brushD, paths := range strokes {
		scale := min(math.Sqrt(t.I.X*t.I.X+t.I.Y*t.I.Y) + math.Sqrt(t.J.X*t.J.X+t.J.Y*t.J.Y))

		newBrushD := scale * brushD
		res[newBrushD] = append(res[newBrushD], paths.Transform(t)...)
	}

	return res
}
