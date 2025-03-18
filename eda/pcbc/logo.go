// Copyright © 2025 Alex Temnok. All rights reserved.

package pcbc

import (
	"strconv"
	"strings"
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Logo = &eda.Component{
		Marks: path.Join(
			font.Centered("pc").Apply(transform.Move(0, 0.33)),
			font.Centered("bc").Apply(transform.Move(0, -0.33)),
		).Apply(transform.RotateDegrees(45)),
	}

	Firm = eda.CenteredTextColumn(-1,
		"TMNK",
		"TECH",
	)
)

func Rev(y, m, d int) *eda.Component {
	const base = 36
	ymd := strconv.FormatInt(int64(y)%100, base) +
		strconv.FormatInt(int64(m), base) +
		strconv.FormatInt(int64(d), base)

	return eda.CenteredText(strings.ToUpper(ymd))
}
