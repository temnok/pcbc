// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/path"
	"temnok/pcbc/shape"
)

var etchBitmapSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Remove Pi13/Si13"},
		Index:    &lbrn.Param{Value: "0"},
		Priority: &lbrn.Param{Value: "0"},
		Negative: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "30"},
		QPulseWidth: &lbrn.Param{Value: "30"},
		Frequency:   &lbrn.Param{Value: "3000000"},

		NumPasses: &lbrn.Param{Value: "4"},
		Speed:     &lbrn.Param{Value: "400"},
		Interval:  &lbrn.Param{Value: "0.01"},
		DPI:       &lbrn.Param{Value: "2540"},

		Angle:            &lbrn.Param{Value: "-90"},
		CrossHatch:       &lbrn.Param{Value: "1"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Clean Cu"},
		Index:    &lbrn.Param{Value: "2"},
		Priority: &lbrn.Param{Value: "2"},
		Negative: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "30"},
		QPulseWidth: &lbrn.Param{Value: "2"},
		Frequency:   &lbrn.Param{Value: "280000"},

		NumPasses: &lbrn.Param{Value: "2"},
		Speed:     &lbrn.Param{Value: "400"},
		Interval:  &lbrn.Param{Value: "0.01"},
		DPI:       &lbrn.Param{Value: "2540"},

		Angle:            &lbrn.Param{Value: "-90"},
		CrossHatch:       &lbrn.Param{Value: "1"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
}

var etchCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Scan",
		Name:     &lbrn.Param{Value: "Clean 1"},
		Index:    &lbrn.Param{Value: "1"},
		Priority: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "5"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		CrossHatch: &lbrn.Param{Value: "1"},
		NumPasses:  &lbrn.Param{Value: "1"},
		Speed:      &lbrn.Param{Value: "800"},
		Interval:   &lbrn.Param{Value: "0.02"},
		DPI:        &lbrn.Param{Value: "1270"},
	},
	{
		Type:     "Cut",
		Name:     &lbrn.Param{Value: "FR4 Cut"},
		Index:    &lbrn.Param{Value: "3"},
		Priority: &lbrn.Param{Value: "3"},

		MaxPower:    &lbrn.Param{Value: "90"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		NumPasses:    &lbrn.Param{Value: "1"},
		GlobalRepeat: &lbrn.Param{Value: "50"},
		Speed:        &lbrn.Param{Value: "100"},
	},
	{
		Type:     "Scan",
		Name:     &lbrn.Param{Value: "Clean 2"},
		Index:    &lbrn.Param{Value: "4"},
		Priority: &lbrn.Param{Value: "4"},

		MaxPower:    &lbrn.Param{Value: "5"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		CrossHatch: &lbrn.Param{Value: "1"},
		NumPasses:  &lbrn.Param{Value: "1"},
		Speed:      &lbrn.Param{Value: "800"},
		Interval:   &lbrn.Param{Value: "0.02"},
		DPI:        &lbrn.Param{Value: "1270"},
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
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: etchBitmapSettings,
		CutSetting:    etchCutSettings,
		Shape: append([]*lbrn.Shape{
			lbrn.NewBitmapShape(0, config.LbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(2, config.LbrnBitmapScale(), bm),
		}, cuts...),
	}

	addCleanPasses(config, p)

	return copper, p.SaveToFile(filename)
}

func removeEtchCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	if component.NoClear {
		return
	}

	t := component.Transform.Multiply(config.BitmapTransform())

	clearWidth := 2 * (component.ClearWidth - config.ExtraCopperWidth)

	// Cuts
	cutBrush := shape.Circle(int((clearWidth / 2) * config.PixelsPerMM))
	cutBrush.ForEachPathsPixel(component.Cuts, t, copper.Set1)

	// Pads
	padBrush := shape.Circle(int(clearWidth * config.PixelsPerMM))
	padBrush.ForEachPathsPixel(component.Pads, t, copper.Set1)

	// Tracks
	trackBrush := shape.Circle(int((component.TrackWidth + clearWidth) * config.PixelsPerMM))
	trackBrush.ForEachPathsPixel(component.Tracks, t, copper.Set1)
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
}

func addEtchCuts(config *config.Config, component *eda.Component, cuts *[]*lbrn.Shape) {
	t := component.Transform.Multiply(config.LbrnCenterMove())

	for _, cut := range component.Cuts {
		//*cuts = append(*cuts, lbrn.NewPathWithTabs(2, t, cut))
		*cuts = append(*cuts, lbrn.NewPath(3, t, cut))
	}
}

func addCleanPasses(config *config.Config, p *lbrn.LightBurnProject) {
	t := config.LbrnCenterMove()
	boardBounds := path.Rect(config.Width, config.Height)
	p.Shape = append(p.Shape,
		lbrn.NewPath(1, t, boardBounds),
		lbrn.NewPath(4, t, boardBounds),
	)
}
