// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"temnok/pcbc/bezier"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
)

func SaveStencilGerber(config *config.Config, component *eda.Component) error {
	buf := &bytes.Buffer{}
	buf.WriteString(`G75*
%FSLAX44Y44*%
%MOMM*%
`)

	component.Visit(func(c *eda.Component) {
		if c.OuterCut {
			return
		}

		for _, pad := range c.Pads {
			xy := pad.Transform(c.Transform).ToXY()
			if len(xy) < 8 {
				continue
			}

			x0, y0 := toGerberCoord(xy[0]), toGerberCoord(xy[1])

			fmt.Fprintf(buf, "\nG36*\n")
			fmt.Fprintf(buf, "X%vY%vD2*\n", x0, y0)

			bezier.Linearize(xy, config.StencilLinearizeDelta, func(x1, y1 float64) {
				x, y := toGerberCoord(x1), toGerberCoord(y1)
				if x != x0 {
					fmt.Fprintf(buf, "X%v", x)
				}
				if y != y0 {
					fmt.Fprintf(buf, "Y%v", y)
				}
				if x == x0 && y == y0 {
					return
				}

				x0, y0 = x, y
				fmt.Fprintf(buf, "D1*\n")
			})

			fmt.Fprintf(buf, "G37*\n")
		}
	})

	fmt.Fprintf(buf, "\nM02*\n")

	return os.WriteFile(config.SavePath+"stencil.gtp", buf.Bytes(), 0770)
}

func toGerberCoord(val float64) int {
	return int(math.Round(val * 10_000))
}
