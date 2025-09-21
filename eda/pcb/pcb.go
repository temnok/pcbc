// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"fmt"
	"strings"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/util"
	"temnok/pcbc/util/ptr"
)

func Process(config *config.Config, components ...*eda.Component) error {
	var jobs []func() error

	for _, comp := range components {
		jobs = append(jobs, func() error { return processComponent(config, comp) })
	}

	return util.RunConcurrently(jobs...)
}

func processComponent(initialConfig *config.Config, initialComponent *eda.Component) error {
	if initialConfig == nil {
		initialConfig = config.Default()
	}

	config := *initialConfig
	config.SavePath = strings.ReplaceAll(initialConfig.SavePath, "{}", fmt.Sprint(initialComponent.Layer))

	component := &eda.Component{
		Layer:       initialComponent.Layer,
		TracksWidth: ptr.To(config.TrackWidth),
		ClearWidth:  ptr.To(config.ClearWidth),
		Inner: eda.Components{
			initialComponent,
		},
	}

	var copper, mask, silk, stencil *bitmap.Bitmap

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
			var e error
			stencil, e = SaveStencil(&config, component)
			return e
		},
	)
	if err != nil {
		return err
	}

	return saveOverview(&config, copper, mask, silk, stencil)
}
