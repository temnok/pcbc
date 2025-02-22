// Copyright © 2025 Alex Temnok. All rights reserved.

package path

func Linear(points []Point) Path {
	if len(points) == 0 {
		return nil
	}

	path := Path{points[0]}

	for _, p := range points[1:] {
		path = append(path, path[len(path)-1], p, p)
	}

	return path
}
