// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

var etchBitmapSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Etch"},
		Index:    lbrn.Param{Value: "0"},
		Priority: lbrn.Param{Value: "0"},

		MaxPower:    lbrn.Param{Value: "20"},
		QPulseWidth: lbrn.Param{Value: "200"},
		Frequency:   lbrn.Param{Value: "20000"},

		Speed:            lbrn.Param{Value: "600"},
		Interval:         lbrn.Param{Value: "0.01"},
		DPI:              lbrn.Param{Value: "2540"},
		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.05"},

		Negative: lbrn.Param{Value: "1"},
	},
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Clean 1"},
		Index:    lbrn.Param{Value: "1"},
		Priority: lbrn.Param{Value: "1"},

		MaxPower:    lbrn.Param{Value: "50"},
		QPulseWidth: lbrn.Param{Value: "2"},
		Frequency:   lbrn.Param{Value: "280000"},

		Speed:            lbrn.Param{Value: "2000"},
		Interval:         lbrn.Param{Value: "0.01"},
		DPI:              lbrn.Param{Value: "2540"},
		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.15"},

		Negative: lbrn.Param{Value: "1"},
	},
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Clean 2"},
		Index:    lbrn.Param{Value: "3"},
		Priority: lbrn.Param{Value: "3"},

		MaxPower:    lbrn.Param{Value: "50"},
		QPulseWidth: lbrn.Param{Value: "2"},
		Frequency:   lbrn.Param{Value: "280000"},

		Speed:            lbrn.Param{Value: "2000"},
		Interval:         lbrn.Param{Value: "0.01"},
		DPI:              lbrn.Param{Value: "2540"},
		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.15"},

		Angle:    lbrn.Param{Value: "90"},
		Negative: lbrn.Param{Value: "1"},
	},
}

var etchCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Cut",
		Name:     lbrn.Param{Value: "FR4 Cut"},
		Index:    lbrn.Param{Value: "2"},
		Priority: lbrn.Param{Value: "2"},

		Speed:        lbrn.Param{Value: "100"},
		GlobalRepeat: lbrn.Param{Value: "50"},

		MaxPower:    lbrn.Param{Value: "90"},
		QPulseWidth: lbrn.Param{Value: "200"},
		Frequency:   lbrn.Param{Value: "20000"},

		TabsEnabled: lbrn.Param{Value: "1"},
		TabSize:     lbrn.Param{Value: "0.1"},
	},
}

func SaveEtch(config *config.Config, component *eda.Component) (*bitmap.Bitmap, error) {
	copper := bitmap.New(config.BitmapSizeInPixels())
	var cuts []*lbrn.Shape

	component.Visit(func(c *eda.Component) {
		removeEtchCopper(config, c, copper)
	})

	component.Visit(func(c *eda.Component) {
		addEtchCopper(config, c, copper)
		addEtchCuts(config, c, &cuts)
	})

	filename := config.SavePath + "etch.lbrn"
	im := image.NewSingle(copper, color.Black, color.White)
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		CutSettingImg: etchBitmapSettings,
		CutSetting:    etchCutSettings,
		Shape: append([]*lbrn.Shape{
			lbrn.NewBitmapShape(0, config.LbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(1, config.LbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(3, config.LbrnBitmapScale(), bm),
		}, cuts...),
	}

	return copper, p.SaveToFile(filename)
}

func removeEtchCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.BitmapTransform())

	clearWidth := 2 * (config.CopperClearWidth - config.ExtraCopperWidth)

	// Pads
	padClearBrush := shape.Circle(int(clearWidth * config.PixelsPerMM))
	padClearBrush.ForEachPathsPixel(component.Pads, t, copper.Set1)

	// Non-ground tracks
	brush := shape.Circle(int((component.TrackWidth + clearWidth) * config.PixelsPerMM))
	brush.ForEachPathsPixel(component.Tracks, t, copper.Set1)

	clearBrush := shape.Circle(int(clearWidth * config.PixelsPerMM))
	clearBrush = shape.Circle(int(config.CopperClearWidth * config.PixelsPerMM))
	clearBrush.ForEachPathsPixel(component.Cuts, t, copper.Set1)
	clearBrush.ForEachPathsPixel(component.Perforations, t, copper.Set1)
}

func addEtchCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.BitmapTransform())

	// Pads
	shape.ForEachRow(component.Pads, t, copper.Set0)

	extraCopperBrush := shape.Circle(int(config.ExtraCopperWidth * config.PixelsPerMM))
	extraCopperBrush.ForEachPathsPixel(component.Pads, t, copper.Set0)

	// Tracks
	brush := shape.Circle(int((component.TrackWidth + config.ExtraCopperWidth) * config.PixelsPerMM))
	brush.ForEachPathsPixel(component.Tracks, t, copper.Set0)
	brush.ForEachPathsPixel(component.GroundTracks, t, copper.Set0)
}

func addEtchCuts(config *config.Config, component *eda.Component, cuts *[]*lbrn.Shape) {
	t := component.Transform.Multiply(config.LbrnCenterMove())

	for _, cut := range component.Cuts {
		//*cuts = append(*cuts, lbrn.NewPathWithTabs(2, t, cut))
		*cuts = append(*cuts, lbrn.NewPath(2, t, cut))
	}

	for _, perforation := range component.Perforations {
		*cuts = append(*cuts, lbrn.NewPath(2, t, perforation))
	}
}
