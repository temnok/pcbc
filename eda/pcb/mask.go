// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"errors"
	"image/color"
	"math"
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
	maskSilkPassIndex = 0
	maskCutPassIndex  = 1
)

var maskCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Silk"},
		Index:    &lbrn.Param{Value: strconv.Itoa(maskSilkPassIndex)},
		Priority: &lbrn.Param{Value: strconv.Itoa(maskSilkPassIndex)},

		MaxPower:    &lbrn.Param{Value: "5"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		NumPasses: &lbrn.Param{Value: "1"},
		Speed:     &lbrn.Param{Value: "800"},
		Interval:  &lbrn.Param{Value: "0.02"},
		DPI:       &lbrn.Param{Value: "1270"},

		CrossHatch: &lbrn.Param{Value: "1"},
		Angle:      &lbrn.Param{Value: "-90"},

		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Mask"},
		Index:    &lbrn.Param{Value: strconv.Itoa(maskCutPassIndex)},
		Priority: &lbrn.Param{Value: strconv.Itoa(maskCutPassIndex)},

		MaxPower:    &lbrn.Param{Value: "10"},
		QPulseWidth: &lbrn.Param{Value: "80"},
		Frequency:   &lbrn.Param{Value: "2000000"},

		NumPasses:        &lbrn.Param{Value: "10"},
		Speed:            &lbrn.Param{Value: "500"},
		Interval:         &lbrn.Param{Value: "0.01"},
		DPI:              &lbrn.Param{Value: "2540"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},

		CrossHatch: &lbrn.Param{Value: "1"},
		Angle:      &lbrn.Param{Value: "90"},
	},
}

func SaveMask(config *config.Config, root *eda.Component, back bool) (*bitmap.Bitmap, *bitmap.Bitmap, error) {
	mask := bitmap.New(config.BitmapSizeInPixels())
	silk := bitmap.New(config.BitmapSizeInPixels())
	cuts := bitmap.New(config.BitmapSizeInPixels())

	root.Visit(func(c *eda.Component) {
		addMaskPads(config, c, back, mask)
		addMaskMarks(config, c, back, silk)
		addMaskCuts(config, c, cuts)
	})

	shrunkCuts := func(c *eda.Component) path.Paths {
		return c.AlignCuts
	}
	renderShrunkCuts(config, root, shrunkCuts, mask)

	silk.Or(cuts)

	maskFilename := config.SavePath + fileNamePrefix[back] + "mask.lbrn"
	silkImage := image.NewSingle(silk, color.Transparent, color.Black)
	maskImage := image.NewSingle(mask, color.Transparent, color.Black)

	cutsFilename := config.SavePath + fileNamePrefix[back] + "mask-cut.lbrn"
	cutsImage := image.NewSingle(cuts, color.Transparent, color.Black)

	maskProject := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(maskSilkPassIndex, config.LbrnBitmapScale(), silkImage),
			lbrn.NewBitmapShapeFromImage(maskCutPassIndex, config.LbrnBitmapScale(), maskImage),
		},
	}

	cutsProject := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(maskCutPassIndex, config.LbrnBitmapScale(), cutsImage),
		},
	}

	return mask, silk, errors.Join(
		maskProject.SaveToFile(maskFilename),
		cutsProject.SaveToFile(cutsFilename),
	)
}

func addMaskPads(config *config.Config, c *eda.Component, back bool, mask *bitmap.Bitmap) {
	if c.CutsHidden() || back {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())
	brush := shape.Circle(int(c.CutsWidth * config.PixelsPerMM))
	brush.ForEachPathsPixel(c.Pads, t, mask.Set1)
}

func addMaskMarks(config *config.Config, c *eda.Component, back bool, silk *bitmap.Bitmap) {
	if c.Back != back {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())

	brushW := c.MarksWidth * math.Sqrt(math.Abs(t.Det()))

	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, silk.Set1)
}

func addMaskCuts(config *config.Config, c *eda.Component, cuts *bitmap.Bitmap) {
	if c.CutsHidden() {
		return
	}

	brush := shape.Circle(int(c.ClearWidth * config.PixelsPerMM))
	t := c.Transform.Multiply(config.BitmapTransform())
	brush.ForEachPathsPixel(c.Cuts, t, cuts.Set1)
}
