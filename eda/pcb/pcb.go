// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

type Config struct {
	Width, Height float64
	PixelsPerMM   float64

	TrackWidth       float64
	ExtraCopperWidth float64
	CopperClearWidth float64
	MaskCutWidth     float64
	OverviewCutWidth float64

	LbrnCenterX, LbrnCenterY float64

	SavePath string
}

func Defaults() *Config {
	return &Config{
		PixelsPerMM: 100,

		TrackWidth:       0.25,
		ExtraCopperWidth: 0.05,
		CopperClearWidth: 0.25,
		MaskCutWidth:     0.1,
		OverviewCutWidth: 0.02,

		LbrnCenterX: 55,
		LbrnCenterY: 55,

		SavePath: "out/",
	}
}

func Generate(component *eda.Component) error {
	return SaveFiles(Defaults(), component)
}

func (config *Config) bitmapSize() (int, int) {
	return int(config.Width * config.PixelsPerMM), int(config.Height * config.PixelsPerMM)
}

func (config *Config) bitmapTransform() transform.T {
	return transform.Move(config.Width/2, config.Height/2).ScaleUniformly(config.PixelsPerMM)
}

func (config *Config) lbrnCenterMove() transform.T {
	return transform.Move(config.LbrnCenterX, config.LbrnCenterY)
}

func (config *Config) lbrnBitmapScale() transform.T {
	scale := 1.0 / config.PixelsPerMM
	return transform.Scale(scale, -scale).Multiply(config.lbrnCenterMove())
}

func SaveFiles(initialConfig *Config, component *eda.Component) error {
	config := *initialConfig
	if config.Width == 0 || config.Height == 0 {
		w, h := component.Size()
		w, h = w+1, h+1
		if config.Width == 0 {
			config.Width = w
		}
		if config.Height == 0 {
			config.Height = h
		}
	}

	if component.TrackWidth == 0 {
		component.TrackWidth = config.TrackWidth
	}

	var copper, mask, silk *bitmap.Bitmap

	err := util.RunConcurrently(
		func() error {
			var e error
			copper, e = SaveEtch(&config, component)
			return e
		},
		func() error {
			var e error
			mask, silk, e = SaveMask(&config, component)
			return e
		},
		func() error {
			return SaveStencil(&config, component)
		},
	)
	if err != nil {
		return err
	}

	return SaveOverview(&config, component, copper, mask, silk)
}
