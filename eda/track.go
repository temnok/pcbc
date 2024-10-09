package eda

import (
	"math"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

type Track path.Points

func TrackPaths(lines ...Track) Paths {
	res := make(Paths, len(lines))

	for i, l := range lines {
		res[i] = path.Lines(path.Points(l))
	}

	return res
}

func (track Track) Transform(t geom.Transform) Track {
	return Track(Path(track).Transform(t))
}

func (track Track) X(x float64) Track {
	n := len(track)

	switch {
	case n == 0:
		return append(track, geom.XY{x, 0})
	case n == 1:
		return append(track, geom.XY{x, track[0].Y})
	}

	track = append(track, geom.XY{x, track[n-1].Y})

	if track[n-2].X == track[n-1].X {
		dx := track[n].X - track[n-1].X
		dy := track[n-1].Y - track[n-2].Y

		if ax, ay := math.Abs(dx), math.Abs(dy); ax <= ay {
			track[n-1].Y -= sign(dy) * ax
		} else {
			track[n-1].X += sign(dx) * ay
		}
	}

	return track
}

func (track Track) Y(y float64) Track {
	n := len(track)

	switch {
	case n == 0:
		return append(track, geom.XY{0, y})
	case n == 1:
		return append(track, geom.XY{track[0].X, y})
	}

	track = append(track, geom.XY{track[n-1].X, y})

	if track[n-2].Y == track[n-1].Y {
		dy := track[n].Y - track[n-1].Y
		dx := track[n-1].X - track[n-2].X

		if ay, ax := math.Abs(dy), math.Abs(dx); ay <= ax {
			track[n-1].X -= sign(dx) * ay
		} else {
			track[n-1].Y += sign(dy) * ax
		}
	}

	return track
}

func sign(val float64) float64 {
	switch {
	case val < 0:
		return -1
	case val > 0:
		return 1
	default:
		return 0
	}
}
