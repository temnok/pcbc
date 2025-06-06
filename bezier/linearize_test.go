// Copyright © 2025 Alex Temnok. All rights reserved.

package bezier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearize(t *testing.T) {
	path := circle(10)

	assert.Equal(t, []string{
		"7.071", "7.071",
		"0.000", "10.000",
		"-7.071", "7.071",
		"-10.000", "0.000",
		"-7.071", "-7.071",
		"0.000", "-10.000",
		"7.071", "-7.071",
		"10.000", "0.000",
	}, linearizeToArray(path, 1))
}

func TestLinearizeCurve(t *testing.T) {
	arc := circle(10)[:8]

	assert.Equal(t, []string{
		"7.071", "7.071",
		"0.000", "10.000",
	}, linearizeToArray(arc, 1))

	assert.Equal(t, []string{
		"9.214", "3.892",
		"7.071", "7.071",
		"3.892", "9.214",
		"0.000", "10.000",
	}, linearizeToArray(arc, 0.2))

	assert.Equal(t, []string{
		"9.797", "2.015",
		"9.214", "3.892",
		"8.292", "5.591",
		"7.071", "7.071",
		"5.591", "8.292",
		"3.892", "9.214",
		"2.015", "9.797",
		"0.000", "10.000",
	}, linearizeToArray(arc, 0.1))
}

func linearizeToArray(xy []float64, delta float64) []string {
	var out []string

	Linearize(xy, delta, func(x, y float64) {
		out = append(out, fmt.Sprintf("%.3f", x), fmt.Sprintf("%.3f", y))
	})

	return out
}
