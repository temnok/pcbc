// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

var fileNamePrefix = map[bool]string{
	false: "1-",
	true:  "2-",
}

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

	component := &eda.Component{
		TracksWidth: config.TrackWidth,
		ClearWidth:  config.ClearWidth,
		Nested: eda.Components{
			initialComponent,
		},
	}

	componentBack := component.Arrange(transform.MirrorX())

	var copper1, copper2, mask1, mask2, silk1, silk2, stencil *bitmap.Bitmap

	err := util.RunConcurrently(
		func() error {
			var e error
			copper1, e = SaveEtch(&config, component, false)
			return e
		},
		func() error {
			var e error
			copper2, e = SaveEtch(&config, componentBack, true)
			return e
		},
		func() error {
			var e error
			mask1, silk1, e = SaveMask(&config, component, false)
			return e
		},
		func() error {
			var e error
			mask2, silk2, e = SaveMask(&config, componentBack, true)
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

	return util.RunConcurrently(
		func() error {
			return saveOverview(&config, false, copper1, mask1, silk1, stencil)
		},

		func() error {
			return saveOverview(&config, true, copper2, mask2, silk2, nil)
		},
	)
}
