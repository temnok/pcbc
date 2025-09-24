// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import "temnok/pcbc/font"

func CenteredText(line string) *Component {
	return CenteredTextColumn(0, line)
}

func CenteredTextColumn(dy float64, lines ...string) *Component {
	return &Component{
		Marks: font.CenteredColumn(dy, lines...),
	}
}
