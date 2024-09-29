package mph100imp40f

import (
	"temnok/lab/eda/lib"
	"temnok/lab/path"
)

func Gvsp(n int) *lib.Component {
	const tenth = 2.54

	return &lib.Component{
		Pads:   path.Circle(0.75).Clone(n, tenth, 0),
		Placer: path.Rect(float64(n)*tenth+0.1, tenth+0.1),
	}
}
