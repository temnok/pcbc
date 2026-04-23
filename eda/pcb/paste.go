// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"github.com/temnok/pcbc/bitmap"
	"github.com/temnok/pcbc/bitmap/image"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/lbrn"
	"github.com/temnok/pcbc/shape"
	"image/color"
	"strconv"
)

var pasteBitmapSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Paste"},
		Index:    &lbrn.Param{Value: strconv.Itoa(1)},
		Priority: &lbrn.Param{Value: strconv.Itoa(1)},

		MaxPower:    &lbrn.Param{Value: "5"},
		QPulseWidth: &lbrn.Param{Value: "80"},
		Frequency:   &lbrn.Param{Value: "40000"},

		NumPasses:        &lbrn.Param{Value: "5"},
		Speed:            &lbrn.Param{Value: "500"},
		Interval:         &lbrn.Param{Value: "0.01"},
		DPI:              &lbrn.Param{Value: "2540"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},

		CrossHatch: &lbrn.Param{Value: "1"},
		Angle:      &lbrn.Param{Value: "90"},
	},
}

func SavePaste(config *config.Config, component *eda.Component) (*bitmap.Bitmap, error) {
	paste := bitmap.New(config.BitmapSizeInPixels())

	component.Visit(func(c *eda.Component) {
		addPaste(config, c, paste)
	})

	filename := config.SavePath + "0-paste.lbrn"
	pasteImage := image.NewSingle(paste, color.Transparent, color.Black)
	pasteBitmap := lbrn.NewBase64Bitmap(pasteImage)

	p := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: pasteBitmapSettings,
		Shape: append([]*lbrn.Shape{
			lbrn.NewBitmapShape(1, config.LbrnBitmapScale(), pasteBitmap),
		}),
	}

	return paste, p.SaveToFile(filename)
}

func addPaste(config *config.Config, c *eda.Component, paste *bitmap.Bitmap) {
	t := c.Transform.Multiply(config.BitmapTransform())

	// Pads
	shape.ForEachRow(c.Pads, t, paste.Set1)

	// Tracks
	brush := shape.Circle(int(c.TracksWidth * config.PixelsPerMM))
	brush.ForEachPathsPixel(c.Tracks, t, paste.Set1)
}
