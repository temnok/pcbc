// Copyright © 2025 Alex Temnok. All rights reserved.

package config

import (
	"temnok/pcbc/transform"
)

type Config struct {
	Width, Height float64
	PixelsPerMM   float64

	TrackWidth float64
	ClearWidth float64

	ExtraCopperWidth      float64
	MaskCutWidth          float64
	MaskPerforationStep   float64
	StencilPadDist        float64
	StencilLinearizeDelta float64

	LbrnCenterX, LbrnCenterY float64

	SavePath string
}

func Default() *Config {
	return &Config{
		PixelsPerMM: 100,

		TrackWidth:       0.2,
		ClearWidth:       0.2,
		ExtraCopperWidth: 0.05,

		MaskCutWidth:          0.1,
		MaskPerforationStep:   0.175,
		StencilPadDist:        0.025,
		StencilLinearizeDelta: 0.01,

		LbrnCenterX: 55,
		LbrnCenterY: 55,

		SavePath: "out/",
	}
}

func (config *Config) BitmapSizeInPixels() (int, int) {
	return int(config.Width * config.PixelsPerMM), int(config.Height * config.PixelsPerMM)
}

func (config *Config) BitmapTransform() transform.T {
	return transform.Move(config.Width/2, config.Height/2).ScaleUniformly(config.PixelsPerMM)
}

func (config *Config) LbrnCenterMove() transform.T {
	return transform.Move(config.LbrnCenterX, config.LbrnCenterY)
}

func (config *Config) LbrnBitmapScale() transform.T {
	scale := 1.0 / config.PixelsPerMM
	return transform.Scale(scale, -scale).Multiply(config.LbrnCenterMove())
}
