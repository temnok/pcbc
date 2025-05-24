// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

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
			font.Centered("pc").Transform(transform.Move(0, 0.3)),
			font.Centered("bc").Transform(transform.Move(0, -0.3)),
		).Transform(transform.RotateDegrees(45)),
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
