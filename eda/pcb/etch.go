// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"strconv"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/path"
	"temnok/pcbc/shape"
)

const (
	etchPassIndex  = 1
	cutPassIndex   = 2
	cleanPassIndex = 3
)

var etchBitmapSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Etch"},
		Index:    &lbrn.Param{Value: strconv.Itoa(etchPassIndex)},
		Priority: &lbrn.Param{Value: strconv.Itoa(etchPassIndex)},

		MaxPower:    &lbrn.Param{Value: "35"},
		QPulseWidth: &lbrn.Param{Value: "80"},
		Frequency:   &lbrn.Param{Value: "40000"},

		NumPasses: &lbrn.Param{Value: "15"},
		Speed:     &lbrn.Param{Value: "800"},
		Interval:  &lbrn.Param{Value: "0.02"},
		DPI:       &lbrn.Param{Value: "1270"},

		Angle:            &lbrn.Param{Value: "-90"},
		CrossHatch:       &lbrn.Param{Value: "1"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
}

func etchCutSettings(bottom bool) []*lbrn.CutSetting {
	var doOutput *lbrn.Param
	if bottom {
		doOutput = &lbrn.Param{Value: "0"}
	}

	return []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     &lbrn.Param{Value: "Cut"},
			Index:    &lbrn.Param{Value: strconv.Itoa(cutPassIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(cutPassIndex)},
			DoOutput: doOutput,

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "200"},
			Frequency:   &lbrn.Param{Value: "20000"},

			NumPasses:    &lbrn.Param{Value: "1"},
			GlobalRepeat: &lbrn.Param{Value: "200"},
			Speed:        &lbrn.Param{Value: "700"},

			SubLayer: &lbrn.SubLayer{
				Type:  "Cut",
				Index: "1",

				MaxPower: &lbrn.Param{Value: "0.1"},
				Speed:    &lbrn.Param{Value: "140"},

				QPulseWidth: &lbrn.Param{Value: "200"},
				Frequency:   &lbrn.Param{Value: "20000"},
			},
		},
		{
			Type:     "Scan",
			Name:     &lbrn.Param{Value: "Clean"},
			Index:    &lbrn.Param{Value: strconv.Itoa(cleanPassIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(cleanPassIndex)},

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "1"},
			Frequency:   &lbrn.Param{Value: "650000"},

			CrossHatch: &lbrn.Param{Value: "1"},
			NumPasses:  &lbrn.Param{Value: "1"},
			Speed:      &lbrn.Param{Value: "6500"},
			Interval:   &lbrn.Param{Value: "0.01"},
			DPI:        &lbrn.Param{Value: "2540"},
		},
	}
}

func SaveEtch(config *config.Config, component *eda.Component, bottom bool) (*bitmap.Bitmap, error) {
	copper := bitmap.New(config.BitmapSizeInPixels())
	var cuts []*lbrn.Shape

	component.Visit(func(c *eda.Component) {
		removeCopper(config, c, bottom, copper)
	})

	component.Visit(func(c *eda.Component) {
		addCopper(config, c, bottom, copper)
	})

	component.Visit(func(c *eda.Component) {
		addCuts(config, c, bottom, copper, &cuts)
	})

	filename := config.SavePath + fileNamePrefix[bottom] + "etch.lbrn"
	copperImage := image.NewSingle(copper, color.Transparent, color.Black)
	copperBitmap := lbrn.NewBase64Bitmap(copperImage)

	p := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: etchBitmapSettings,
		CutSetting:    etchCutSettings(bottom),
		Shape: append([]*lbrn.Shape{
			lbrn.NewBitmapShape(etchPassIndex, config.LbrnBitmapScale(), copperBitmap),
		}, cuts...),
	}

	addCleanPasses(config, p)

	return copper, p.SaveToFile(filename)
}

func removeCopper(config *config.Config, c *eda.Component, bottom bool, copper *bitmap.Bitmap) {
	if c.ClearOff() || c.Bottom != bottom {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())

	// Pads
	padBrush := shape.Circle(int(2 * c.ClearWidth * config.PixelsPerMM))
	padBrush.ForEachPathsPixel(c.Pads, t, copper.Set1)

	// Tracks
	trackBrush := shape.Circle(int((c.TracksWidth + 2*c.ClearWidth) * config.PixelsPerMM))
	trackBrush.ForEachPathsPixel(c.Tracks, t, copper.Set1)
}

func addCopper(config *config.Config, c *eda.Component, bottom bool, copper *bitmap.Bitmap) {
	if c.Bottom != bottom {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())

	// Pads
	shape.ForEachRow(c.Pads, t, copper.Set0)

	// Tracks
	brush := shape.Circle(int(c.TracksWidth * config.PixelsPerMM))
	brush.ForEachPathsPixel(c.Tracks, t, copper.Set0)
}

func addCuts(config *config.Config, c *eda.Component, bottom bool, copper *bitmap.Bitmap, cuts *[]*lbrn.Shape) {
	t := c.Transform.Multiply(config.LbrnCenterMove())

	if !c.Bottom {
		for _, cut := range c.AlignCuts {
			*cuts = append(*cuts, lbrn.NewPath(cutPassIndex, t, cut))
		}
	}

	for _, cut := range c.Cuts {
		*cuts = append(*cuts, lbrn.NewPath(cutPassIndex, t, cut))
	}

	if !bottom {
		t := c.Transform.Multiply(config.BitmapTransform())

		cutBrush := shape.Circle(int(c.ClearWidth * config.PixelsPerMM))

		if !c.Bottom {
			cutBrush.ForEachPathsPixel(c.AlignCuts, t, copper.Set1)
		}

		cutBrush.ForEachPathsPixel(c.Cuts, t, copper.Set1)
	}
}

func addCleanPasses(config *config.Config, p *lbrn.LightBurnProject) {
	t := config.LbrnCenterMove()
	boardBounds := path.Rect(config.Width, config.Height)
	p.Shape = append(p.Shape,
		lbrn.NewPath(cleanPassIndex, t, boardBounds),
	)
}
