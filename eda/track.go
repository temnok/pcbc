// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"math"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type Track []path.Point

func Tracks(tracks ...Track) path.Paths {
	res := make(path.Paths, len(tracks))

	for i, track := range tracks {
		res[i] = path.Linear(track)
	}

	return res
}

func (track Track) Apply(t transform.T) Track {
	return Track(path.Path(track).Apply(t))
}

func (track Track) X(x float64) Track {
	n := len(track)

	switch {
	case n == 0:
		return append(track, path.Point{X: x})
	case n == 1:
		return append(track, path.Point{X: x, Y: track[0].Y})
	}

	track = append(track, path.Point{X: x, Y: track[n-1].Y})

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
		return append(track, path.Point{Y: y})
	case n == 1:
		return append(track, path.Point{X: track[0].X, Y: y})
	}

	track = append(track, path.Point{X: track[n-1].X, Y: y})

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

func (track Track) DX(dx float64) Track {
	return track.X(track[len(track)-1].X + dx)
}

func (track Track) DY(dy float64) Track {
	return track.Y(track[len(track)-1].Y + dy)
}

func (track Track) XY(p path.Point) Track {
	return track.X(p.X).Y(p.Y)
}

func (track Track) YX(p path.Point) Track {
	return track.Y(p.Y).X(p.X)
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
