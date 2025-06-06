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

			x0, y0 := xy[0], xy[1]

			fmt.Fprintf(buf, "\nG36*\n")
			fmt.Fprintf(buf, "X%vY%vD2*\n", toGerberDecimal(x0), toGerberDecimal(y0))

			bezier.Linearize(xy, config.StencilLinearizeDelta, func(x, y float64) {
				fmt.Fprintf(buf, "X%vY%vD1*\n", toGerberDecimal(x), toGerberDecimal(y))
			})

			fmt.Fprintf(buf, "G37*\n")
		}
	})

	fmt.Fprintf(buf, "\nM02*\n")

	return os.WriteFile(config.SavePath+"stencil.gtp", buf.Bytes(), 0770)
}

func toGerberDecimal(val float64) int {
	return int(math.Round(val * 10_000))
}
